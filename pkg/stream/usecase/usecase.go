package usecase

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtspv2"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
	"github.com/mikerumy/vhosting/internal/models"
)

const (
	errorStreamExitNoViewer = "stream exit on demand - no Viewer"
	videoTimeoutSeconds     = 80
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
	fmt.Println("    - ServeStreams()")

	for key, val := range u.cfg.Streams {
		if !val.OnDemand {
			go u.rtspWorkerLoop(key, val.URL, val.OnDemand, val.DisableAudio, val.Debug)
		}
	}
}

func (u *StreamUseCase) rtspWorkerLoop(name, url string, onDemand, disableAudio, debug bool) {
	fmt.Println("    - rtspWorkerLoop(name, url string, onDemand, disableAudio, debug bool)")

	defer u.runUnlock(name)
	for {
		fmt.Println("      info: stream tries to connect", name)
		err := u.rtspWorker(name, url, onDemand, disableAudio, debug)
		if err != nil {
			fmt.Println("      error: rtspWorker. error:", err.Error())
			u.cfg.LastError = err
		}
		if onDemand && !u.isHasViewer(name) {
			fmt.Println("      error:", errorStreamExitNoViewer)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (u *StreamUseCase) rtspWorker(name, url string, onDemand, disableAudio, debug bool) error {
	fmt.Println("    - rtspWorker(name, url string, onDemand, disableAudio, debug bool) error")

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
					return errors.New(errorStreamExitNoViewer)
				} else {
					clientTest.Reset(20 * time.Second)
				}
			}
		case <-keyTest.C:
			return errors.New("stream exit - no video on stream")
		case signals := <-rtspClient.Signals:
			switch signals {
			case rtspv2.SignalCodecUpdate:
				u.codecAdd(name, rtspClient.CodecData)
			case rtspv2.SignalStreamRTPStop:
				return errors.New("stream exit - rtsp disconnect")
			}
		case packetAV := <-rtspClient.OutgoingPacketQueue:
			if audioOnly || packetAV.IsKeyFrame {
				keyTest.Reset(20 * time.Second)
			}
			u.cast(name, *packetAV)
		}
	}
}

func (u *StreamUseCase) codecAdd(suuid string, codecs []av.CodecData) {
	fmt.Println("    - codecAdd(suuid string, codecs []av.CodecData)")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	t := u.cfg.Streams[suuid]
	t.Codecs = codecs
	u.cfg.Streams[suuid] = t
}

func (u *StreamUseCase) isHasViewer(uuid string) bool {
	fmt.Println("    - isHasViewer(uuid string) bool")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	if cfg, ok := u.cfg.Streams[uuid]; ok && len(cfg.ClientList) > 0 {
		return true
	}
	return false
}

func (u *StreamUseCase) cast(uuid string, pck av.Packet) {
	fmt.Println("    - cast(uuid string, pck av.Packet)")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	for _, val := range u.cfg.Streams[uuid].ClientList {
		if len(val.Cast) < cap(val.Cast) {
			val.Cast <- pck
		}
	}
}

func (u *StreamUseCase) runUnlock(uuid string) {
	fmt.Println("    - runUnlock(uuid string)")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	if cfg, ok := u.cfg.Streams[uuid]; ok {
		if cfg.OnDemand && cfg.RunLock {
			cfg.RunLock = false
			u.cfg.Streams[uuid] = cfg
		}
	}
}

func (u *StreamUseCase) Exit(suuid string) bool {
	fmt.Println("    - Exit(suuid string) bool")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	_, ok := u.cfg.Streams[suuid]
	return ok
}

func (u *StreamUseCase) RunIfNotRun(uuid string) {
	fmt.Println("    - RunIfNotRun(uuid string)")

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

func (u *StreamUseCase) CodecGet(suuid string) []av.CodecData {
	fmt.Println("    - CodecGet(suuid string) []av.CodecData")

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
					if codecVideo.SPS() == nil && codecVideo.PPS() == nil &&
						len(codecVideo.SPS()) <= 0 && len(codecVideo.PPS()) <= 0 {
						fmt.Println("      error: bad video codec - waiting for SPS/PPS")
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

func (u *StreamUseCase) GetICEServers() []string {
	fmt.Println("    - GetICEServers() []string")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.ICEServers
}

func (u *StreamUseCase) GetICEUsername() string {
	fmt.Println("    - GetICEUsername() string")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.ICEUsername
}

func (u *StreamUseCase) GetICECredential() string {
	fmt.Println("    - GetICECredential() string")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.ICECredential
}

func (u *StreamUseCase) GetWebRTCPortMin() uint16 {
	fmt.Println("    - GetWebRTCPortMin() uint16")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.WebRTCPortMin
}

func (u *StreamUseCase) GetWebRTCPortMax() uint16 {
	fmt.Println("    - GetWebRTCPortMax() uint16")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	return u.cfg.Server.WebRTCPortMax
}

func (u *StreamUseCase) WritePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool) {
	fmt.Println("    - WritePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool)")

	cid, ch := u.CastListAdd(url)
	defer u.CastListDelete(url, cid)
	defer muxerWebRTC.Close()
	videoStart := false
	noVideo := time.NewTimer(videoTimeoutSeconds * time.Second)
	for {
		select {
		case <-noVideo.C:
			fmt.Println("      info: no video")
			return
		case pck := <-ch:
			if pck.IsKeyFrame || audioOnly {
				noVideo.Reset(videoTimeoutSeconds * time.Second)
				videoStart = true
			}
			if !videoStart && !audioOnly {
				continue
			}
			err := muxerWebRTC.WritePacket(pck)
			if err != nil {
				fmt.Println("      error: WritePacket. error:", err.Error())
				return
			}
		}
	}
}

func (u *StreamUseCase) CastListAdd(suuid string) (string, chan av.Packet) {
	fmt.Println("    - CastListAdd(suuid string) (string, chan av.Packet)")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	cuuid := u.pseudoUUID()
	ch := make(chan av.Packet, 100)
	u.cfg.Streams[suuid].ClientList[cuuid] = models.Viewer{Cast: ch}
	return cuuid, ch
}

func (u *StreamUseCase) pseudoUUID() (uuid string) {
	fmt.Println("    - pseudoUUID() (uuid string)")

	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("      error:", err.Error())
		return
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
	return
}

func (u *StreamUseCase) CastListDelete(suuid, cuuid string) {
	fmt.Println("    - CastListDelete(suuid, cuuid string)")

	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	delete(u.cfg.Streams[suuid].ClientList, cuuid)
}

func (u *StreamUseCase) List() (string, []string) {
	fmt.Println("    - List() (string, []string)")

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
