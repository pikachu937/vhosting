package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/deepch/vdk/av"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/stream"
)

const videoTimeoutSeconds = 80

type StreamHandler struct {
	useCase stream.StreamUseCase
	cfg     *models.ConfigST
}

func NewStreamHandler(useCase stream.StreamUseCase, cfg *models.ConfigST) *StreamHandler {
	return &StreamHandler{
		useCase: useCase,
		cfg:     cfg,
	}
}

func (h *StreamHandler) ServeIndex(ctx *gin.Context) {
	fmt.Printf("  * ServeIndex(ctx *gin.Context)\n")
	_, list := h.useCase.List()
	if len(list) > 0 {
		ctx.Header("Cache-Control", "no-cache, max-age=0, must-revalidate, no-store")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Redirect(http.StatusMovedPermanently, "stream/player/"+list[0])
	} else {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"port":    h.cfg.Server.HTTPPort,
			"version": time.Now().String(),
		})
	}
}

func (h *StreamHandler) ServeStreamPlayer(ctx *gin.Context) {
	fmt.Printf("  * ServeStreamPlayer(ctx *gin.Context)\n")
	_, list := h.useCase.List()
	sort.Strings(list)
	ctx.HTML(http.StatusOK, "player.tmpl", gin.H{
		"port":     h.cfg.Server.HTTPPort,
		"suuid":    ctx.Param("uuid"),
		"suuidMap": list,
		"version":  time.Now().String(),
	})
}

func (h *StreamHandler) ServeStreamCodec(ctx *gin.Context) {
	fmt.Printf("  * ServeStreamCodec(ctx *gin.Context)\n")
	if h.useCase.Exit(ctx.Param("uuid")) {
		h.useCase.RunIfNotRun(ctx.Param("uuid"))
		codecs := h.useCase.CodecGet(ctx.Param("uuid"))
		if codecs == nil {
			return
		}
		var tmpCodec []models.JCodec
		for _, codec := range codecs {
			if codec.Type() != av.H264 && codec.Type() != av.PCM_ALAW && codec.Type() != av.PCM_MULAW && codec.Type() != av.OPUS {
				log.Println("Codec Not Supported WebRTC ignore this track", codec.Type())
				continue
			}
			if codec.Type().IsVideo() {
				tmpCodec = append(tmpCodec, models.JCodec{Type: "video"})
			} else {
				tmpCodec = append(tmpCodec, models.JCodec{Type: "audio"})
			}
		}
		b, err := json.Marshal(tmpCodec)
		if err == nil {
			_, err = ctx.Writer.Write(b)
			if err != nil {
				log.Printf("Write Codec Info error. Error: %s.\n", err.Error())
				return
			}
		}
	}
}

// stream video over WebRTC
func (h *StreamHandler) ServeStreamWebRTC(ctx *gin.Context) {
	fmt.Printf("  * ServeStreamWebRTC(ctx *gin.Context)\n")
	if !h.useCase.Exit(ctx.PostForm("suuid")) {
		log.Println("Stream Not Found")
		return
	}
	h.useCase.RunIfNotRun(ctx.PostForm("suuid"))
	codecs := h.useCase.CodecGet(ctx.PostForm("suuid"))
	if codecs == nil {
		log.Println("Stream Codec Not Found")
		return
	}
	var audioOnly bool
	if len(codecs) == 1 && codecs[0].Type().IsAudio() {
		audioOnly = true
	}
	muxerWebRTC := webrtc.NewMuxer(webrtc.Options{ICEServers: h.useCase.GetICEServers(), ICEUsername: h.useCase.GetICEUsername(), ICECredential: h.useCase.GetICECredential(), PortMin: h.useCase.GetWebRTCPortMin(), PortMax: h.useCase.GetWebRTCPortMax()})
	answer, err := muxerWebRTC.WriteHeader(codecs, ctx.PostForm("data"))
	if err != nil {
		log.Printf("WriteHeader error. Error: %s.\n", err.Error())
		return
	}
	_, err = ctx.Writer.Write([]byte(answer))
	if err != nil {
		log.Printf("Cannot write bytes error. Error: %s.\n", err.Error())
		return
	}

	go h.writePackets(ctx.PostForm("suuid"), muxerWebRTC, audioOnly)
}

func (h *StreamHandler) ServeStreamWebRTC2(ctx *gin.Context) {
	fmt.Printf("  * ServeStreamWebRTC2(ctx *gin.Context)\n")
	url := ctx.PostForm("url")
	if _, ok := h.cfg.Streams[url]; !ok {
		h.cfg.Streams[url] = models.Stream{
			URL:        url,
			OnDemand:   true,
			ClientList: make(map[string]models.Viewer),
		}
	}

	h.useCase.RunIfNotRun(url)

	codecs := h.useCase.CodecGet(url)
	if codecs == nil {
		log.Println("Stream Codec Not Found")
		ctx.JSON(500, models.ResponseError{Error: h.cfg.LastError.Error()})
		return
	}

	muxerWebRTC := webrtc.NewMuxer(
		webrtc.Options{
			ICEServers: h.useCase.GetICEServers(),
			PortMin:    h.useCase.GetWebRTCPortMin(),
			PortMax:    h.useCase.GetWebRTCPortMax(),
		},
	)

	sdp64 := ctx.PostForm("sdp64")
	answer, err := muxerWebRTC.WriteHeader(codecs, sdp64)
	if err != nil {
		log.Printf("Muxer WriteHeader error. Error: %s.\n", err.Error())
		ctx.JSON(500, models.ResponseError{Error: err.Error()})
		return
	}

	response := models.Response{
		Sdp64: answer,
	}

	for _, codec := range codecs {
		if codec.Type() != av.H264 &&
			codec.Type() != av.PCM_ALAW &&
			codec.Type() != av.PCM_MULAW &&
			codec.Type() != av.OPUS {
			log.Println("Codec Not Supported WebRTC ignore this track", codec.Type())
			continue
		}
		if codec.Type().IsVideo() {
			response.Tracks = append(response.Tracks, "video")
		} else {
			response.Tracks = append(response.Tracks, "audio")
		}
	}

	ctx.JSON(200, response)

	audioOnly := len(codecs) == 1 && codecs[0].Type().IsAudio()

	go h.writePackets(url, muxerWebRTC, audioOnly)
}

func (h *StreamHandler) writePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool) {
	fmt.Printf("  * writePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool)\n")
	cid, ch := h.useCase.CastListAdd(url)
	defer h.useCase.CastListDelete(url, cid)
	defer muxerWebRTC.Close()
	videoStart := false
	noVideo := time.NewTimer(videoTimeoutSeconds * time.Second)
	for {
		select {
		case <-noVideo.C:
			log.Printf("Info: No Video")
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
				log.Printf("Write packet error. Error: %s.", err.Error())
				return
			}
		}
	}
}
