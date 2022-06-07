package usecase

import (
	"crypto/rand"
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/cgo/ffmpeg"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtspv2"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
	sconfig "github.com/mikerumy/vhosting/pkg/config_stream"
	"github.com/mikerumy/vhosting/pkg/stream"
)

const (
	errorStreamExitNoViewer = "stream exit on demand - no Viewer"
	snapshotPath            = "./media/vhosting/%s/images"
	snapshotName            = "snapshot.jpg"
	snapshotPeriodSeconds   = 60
	videoTimeoutSeconds     = 80
	jpegQuality             = 70
)

type StreamUseCase struct {
	cfg *sconfig.Config
}

func NewStreamUseCase(cfg *sconfig.Config) *StreamUseCase {
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

func (u *StreamUseCase) rtspWorkerLoop(name, url string, onDemand, disableAudio, debug bool) {
	defer u.runUnlock(name)
	for {
		log.Println("info. stream tries to connect", name)
		err := u.rtspWorker(name, url, onDemand, disableAudio, debug)
		if err != nil {
			log.Println("error. rtspWorker error. error:", err.Error())
			u.cfg.LastError = err
		}
		if onDemand && !u.isHasViewer(name) {
			log.Println("error. on demand && not has Viewer error:", errorStreamExitNoViewer)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (u *StreamUseCase) rtspWorker(name, url string, onDemand, disableAudio, debug bool) error {
	keyTest := time.NewTimer(20 * time.Second)
	clientTest := time.NewTimer(20 * time.Second)

	// add next timeout
	newRTSPClient := rtspv2.RTSPClientOptions{
		URL:              url,
		DisableAudio:     disableAudio,
		DialTimeout:      3 * time.Second,
		ReadWriteTimeout: 3 * time.Second,
		Debug:            debug,
	}

	rtspClient, err := rtspv2.Dial(newRTSPClient)
	if err != nil {
		return err
	}
	defer rtspClient.Close()

	if rtspClient.CodecData != nil {
		u.codecAdd(name, rtspClient.CodecData)
	}

	audioOnly := false
	videoIDX := 0
	for i, codec := range rtspClient.CodecData {
		if codec.Type().IsVideo() {
			audioOnly = false
			videoIDX = i
		}
	}

	var frameDecoderSingle *ffmpeg.VideoDecoder
	if !audioOnly {
		frameDecoderSingle, err = ffmpeg.NewVideoDecoder(rtspClient.CodecData[videoIDX].(av.VideoCodecData))
		if err != nil {
			log.Fatalln("fatal. frameDecoderSingle error. error:", err)
		}
	}

	isTimeToSnapshot := true
	go func() {
		for {
			time.Sleep(snapshotPeriodSeconds * time.Second)
			isTimeToSnapshot = true
		}
	}()

	snapshotDir := fmt.Sprintf(snapshotPath, name)
	if exists, _ := isPathExists(snapshotDir); !exists {
		os.MkdirAll(snapshotDir, 0777)
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
			// sample single frame decode encode to jpeg, save on disk
			if !packetAV.IsKeyFrame {
				break
			}
			pic, err := frameDecoderSingle.DecodeSingle(packetAV.Data)
			if err != nil ||
				pic == nil || !isTimeToSnapshot {
				break
			}
			out, err := os.Create(snapshotDir + "/" + snapshotName)
			if err != nil {
				break
			}
			if err := jpeg.Encode(out, &pic.Image, &jpeg.Options{Quality: jpegQuality}); err == nil {
				log.Printf("info. snapshot created for %s\n", name)
				isTimeToSnapshot = false
			}
		}
	}
}

func isPathExists(snapshotPath string) (bool, error) {
	_, err := os.Stat(snapshotPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (u *StreamUseCase) codecAdd(suuid string, codecs []av.CodecData) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	t := u.cfg.Streams[suuid]
	t.Codecs = codecs
	u.cfg.Streams[suuid] = t
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

func (u *StreamUseCase) runUnlock(uuid string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	cfg, ok := u.cfg.Streams[uuid]
	if !ok {
		return
	}
	if cfg.OnDemand && cfg.RunLock {
		cfg.RunLock = false
		u.cfg.Streams[uuid] = cfg
	}
}

func (u *StreamUseCase) Exit(suuid string) bool {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	_, ok := u.cfg.Streams[suuid]
	return ok
}

func (u *StreamUseCase) RunIfNotRun(uuid string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	cfg, ok := u.cfg.Streams[uuid]
	if !ok {
		return
	}
	if cfg.OnDemand && !cfg.RunLock {
		cfg.RunLock = true
		u.cfg.Streams[uuid] = cfg
		go u.rtspWorkerLoop(uuid, cfg.URL, cfg.OnDemand, cfg.DisableAudio, cfg.Debug)
	}
}

func (u *StreamUseCase) CodecGet(suuid string) []av.CodecData {
	for i := 0; i < 100; i++ {
		u.cfg.Mutex.RLock()
		cfg, ok := u.cfg.Streams[suuid]
		u.cfg.Mutex.RUnlock()
		if !ok {
			return nil
		}
		if cfg.Codecs == nil {
			time.Sleep(50 * time.Millisecond)
			continue
		}
		for _, codec := range cfg.Codecs {
			if codec.Type() != av.H264 {
				continue
			}
			codecVideo := codec.(h264parser.CodecData)
			if codecVideo.SPS() == nil && codecVideo.PPS() == nil &&
				len(codecVideo.SPS()) <= 0 && len(codecVideo.PPS()) <= 0 {
				log.Println("error: bad video codec - waiting for SPS/PPS")
				time.Sleep(50 * time.Millisecond)
			}
		}
		return cfg.Codecs
	}
	return nil
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

func (u *StreamUseCase) WritePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool) {
	cid, ch := u.CastListAdd(url)
	defer u.CastListDelete(url, cid)
	defer muxerWebRTC.Close()
	videoStart := false
	noVideo := time.NewTimer(videoTimeoutSeconds * time.Second)
	for {
		select {
		case <-noVideo.C:
			log.Println("info: no video")
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
				log.Println("error: WritePacket error. error:", err.Error())
				return
			}
		}
	}
}

func (u *StreamUseCase) CastListAdd(suuid string) (string, chan av.Packet) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	cuuid := u.pseudoUUID()
	ch := make(chan av.Packet, 100)
	u.cfg.Streams[suuid].ClientList[cuuid] = stream.Viewer{Cast: ch}
	return cuuid, ch
}

func (u *StreamUseCase) pseudoUUID() (uuid string) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("error. pseudoUUID read error. error:", err.Error())
		return
	}
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
	return
}

func (u *StreamUseCase) CastListDelete(suuid, cuuid string) {
	u.cfg.Mutex.Lock()
	defer u.cfg.Mutex.Unlock()
	delete(u.cfg.Streams[suuid].ClientList, cuuid)
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
