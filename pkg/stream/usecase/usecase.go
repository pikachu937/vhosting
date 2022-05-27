package usecase

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtspv2"
	"github.com/mikerumy/vhosting/internal/models"
)

const (
	ErrorStreamExitNoVideoOnStream = "Stream Exit No Video On Stream"
	ErrorStreamExitRtspDisconnect  = "Stream Exit Rtsp Disconnect"
	ErrorStreamExitNoViewer        = "Stream Exit On Demand No Viewer"
)

type StreamUseCase struct {
	cfg *models.ConfigST
}

func NewStreamUseCase(cfg *models.ConfigST) *StreamUseCase {
	return &StreamUseCase{
		cfg: cfg,
	}
}

func (u *StreamUseCase) ServeStreams() {
	for key, val := range u.cfg.Streams {
		if !val.OnDemand {
			go u.rtspWorkerLoop(key, val.URL, val.OnDemand, val.DisableAudio, val.Debug)
		}
	}
}

func (u *StreamUseCase) RunIfNotRun(uuid string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	if cfg, ok := u.cfg.Streams[uuid]; ok {
		if cfg.OnDemand && !cfg.RunLock {
			cfg.RunLock = true
			u.cfg.Streams[uuid] = cfg
			go u.rtspWorkerLoop(uuid, cfg.URL, cfg.OnDemand, cfg.DisableAudio, cfg.Debug)
		}
	}
}

func (u *StreamUseCase) GetICEServers() []string {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.ICEServers
}

func (u *StreamUseCase) GetICEUsername() string {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.ICEUsername
}

func (u *StreamUseCase) GetICECredential() string {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.ICECredential
}

func (u *StreamUseCase) GetWebRTCPortMin() uint16 {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.WebRTCPortMin
}

func (u *StreamUseCase) GetWebRTCPortMax() uint16 {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.WebRTCPortMax
}

func (u *StreamUseCase) List() (string, []string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	var res []string
	var first string
	for key := range u.cfg.Streams {
		if first == "" {
			first = key
		}
		res = append(res, key)
	}
	return first, res
}

func (u *StreamUseCase) Exit(suuid string) bool {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	_, ok := u.cfg.Streams[suuid]
	return ok
}

func (u *StreamUseCase) CodecGet(suuid string) []av.CodecData {
	for i := 0; i < 100; i++ {
		u.cfg.Mutex.RLock()
		cfg, ok := u.cfg.Streams[suuid]
		u.cfg.Mutex.RUnlock()
		if !ok {
			return nil
		}
		if cfg.Codecs != nil {
			for _, codec := range cfg.Codecs {
				if codec.Type() == av.H264 {
					codecVideo := codec.(h264parser.CodecData)
					if codecVideo.SPS() != nil && codecVideo.PPS() != nil && len(codecVideo.SPS()) > 0 && len(codecVideo.PPS()) > 0 {
						// video ready to play
					} else {
						// video codec not ok
						log.Println("Bad Video Codec SPS or PPS Wait")
						time.Sleep(50 * time.Millisecond)
						continue
					}
				}
			}
			return cfg.Codecs
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}

func (u *StreamUseCase) CastListAdd(suuid string) (string, chan av.Packet) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	cuuid := u.pseudoUUID()
	ch := make(chan av.Packet, 100)
	u.cfg.Streams[suuid].ClientList[cuuid] = models.Viewer{Cast: ch}
	return cuuid, ch
}

func (u *StreamUseCase) CastListDelete(suuid, cuuid string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	delete(u.cfg.Streams[suuid].ClientList, cuuid)
}

func (u *StreamUseCase) rtspWorkerLoop(name, url string, onDemand, disableAudio, debug bool) {
	defer u.runUnlock(name)
	for {
		log.Println("Stream Try Connect", name)
		err := u.rtspWorker(name, url, onDemand, disableAudio, debug)
		if err != nil {
			log.Println(err)
			u.cfg.LastError = err
		}
		if onDemand && !u.isHasViewer(name) {
			log.Println(ErrorStreamExitNoViewer)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (u *StreamUseCase) rtspWorker(name, url string, onDemand, disableAudio, debug bool) error {
	keyTest := time.NewTimer(20 * time.Second)
	clientTest := time.NewTimer(20 * time.Second)
	// add next timeout
	rtspClient, err := rtspv2.Dial(rtspv2.RTSPClientOptions{URL: url, DisableAudio: disableAudio, DialTimeout: 3 * time.Second, ReadWriteTimeout: 3 * time.Second, Debug: debug})
	if err != nil {
		return err
	}
	defer rtspClient.Close()
	if rtspClient.CodecData != nil {
		u.codecAdd(name, rtspClient.CodecData)
	}
	var audioOnly bool
	if len(rtspClient.CodecData) == 1 && rtspClient.CodecData[0].Type().IsAudio() {
		audioOnly = true
	}
	for {
		select {
		case <-clientTest.C:
			if onDemand {
				if !u.isHasViewer(name) {
					return errors.New(ErrorStreamExitNoViewer)
				} else {
					clientTest.Reset(20 * time.Second)
				}
			}
		case <-keyTest.C:
			return errors.New(ErrorStreamExitNoVideoOnStream)
		case signals := <-rtspClient.Signals:
			switch signals {
			case rtspv2.SignalCodecUpdate:
				u.codecAdd(name, rtspClient.CodecData)
			case rtspv2.SignalStreamRTPStop:
				return errors.New(ErrorStreamExitRtspDisconnect)
			}
		case packetAV := <-rtspClient.OutgoingPacketQueue:
			if audioOnly || packetAV.IsKeyFrame {
				keyTest.Reset(20 * time.Second)
			}
			u.cast(name, *packetAV)
		}
	}
}

func (u *StreamUseCase) runUnlock(uuid string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	if cfg, ok := u.cfg.Streams[uuid]; ok {
		if cfg.OnDemand && cfg.RunLock {
			cfg.RunLock = false
			u.cfg.Streams[uuid] = cfg
		}
	}
}

func (u *StreamUseCase) isHasViewer(uuid string) bool {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	if cfg, ok := u.cfg.Streams[uuid]; ok && len(cfg.ClientList) > 0 {
		return true
	}
	return false
}

func (u *StreamUseCase) cast(uuid string, pck av.Packet) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	for _, val := range u.cfg.Streams[uuid].ClientList {
		if len(val.Cast) < cap(val.Cast) {
			val.Cast <- pck
		}
	}
}

func (u *StreamUseCase) codecAdd(suuid string, codecs []av.CodecData) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	t := u.cfg.Streams[suuid]
	t.Codecs = codecs
	u.cfg.Streams[suuid] = t
}

func (u *StreamUseCase) pseudoUUID() (uuid string) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
	return
}
