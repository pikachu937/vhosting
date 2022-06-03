package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/deepch/vdk/av"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
	"github.com/gin-gonic/gin"
	sconfig "github.com/mikerumy/vhosting/pkg/config_stream"
	"github.com/mikerumy/vhosting/pkg/stream"
)

type StreamHandler struct {
	cfg     *sconfig.Config
	useCase stream.StreamUseCase
}

func NewStreamHandler(cfg *sconfig.Config, useCase stream.StreamUseCase) *StreamHandler {
	return &StreamHandler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func (h *StreamHandler) ServeIndex(ctx *gin.Context) {
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
	uuid := ctx.Param("uuid")
	if h.useCase.Exit(uuid) {
		h.useCase.RunIfNotRun(uuid)

		codecs := h.useCase.CodecGet(uuid)
		if codecs == nil {
			return
		}

		var tmpCodec []stream.JCodec
		for _, codec := range codecs {
			if codec.Type() != av.H264 && codec.Type() != av.PCM_ALAW && codec.Type() != av.PCM_MULAW && codec.Type() != av.OPUS {
				fmt.Println("    error: track is ignored - codec not supported WebRTC. codec type:", codec.Type())
				continue
			}

			if codec.Type().IsVideo() {
				tmpCodec = append(tmpCodec, stream.JCodec{Type: "video"})
			} else {
				tmpCodec = append(tmpCodec, stream.JCodec{Type: "audio"})
			}
		}

		b, err := json.Marshal(tmpCodec)
		if err == nil {
			_, err = ctx.Writer.Write(b)
			if err != nil {
				fmt.Println("    error: writing of codec info. error:", err.Error())
				return
			}
		}
	}
}

func (h *StreamHandler) ServeStreamVidOverWebRTC(ctx *gin.Context) {
	suuid := ctx.PostForm("suuid")
	if !h.useCase.Exit(suuid) {
		fmt.Println("    info: stream not found. suuid:", suuid)
		return
	}

	h.useCase.RunIfNotRun(suuid)

	codecs := h.useCase.CodecGet(suuid)
	if codecs == nil {
		fmt.Println("    info: stream codec not found. suuid:", suuid)
		return
	}

	audioOnly := false
	if len(codecs) == 1 && codecs[0].Type().IsAudio() {
		audioOnly = true
	}

	muxerWebRTC := webrtc.NewMuxer(webrtc.Options{ICEServers: h.useCase.GetICEServers(),
		ICEUsername: h.useCase.GetICEUsername(), ICECredential: h.useCase.GetICECredential(),
		PortMin: h.useCase.GetWebRTCPortMin(), PortMax: h.useCase.GetWebRTCPortMax()})
	answer, err := muxerWebRTC.WriteHeader(codecs, ctx.PostForm("data"))
	if err != nil {
		fmt.Println("    error: WriteHeader. error:", err.Error())
		return
	}

	if _, err := ctx.Writer.Write([]byte(answer)); err != nil {
		fmt.Println("    error: cannot write bytes. error:", err.Error())
		return
	}

	go h.useCase.WritePackets(suuid, muxerWebRTC, audioOnly)
}

func (h *StreamHandler) ServeStreamWebRTC2(ctx *gin.Context) {
	url := ctx.PostForm("url")
	if _, ok := h.cfg.Streams[url]; !ok {
		h.cfg.Streams[url] = stream.Stream{
			URL:        url,
			OnDemand:   true,
			ClientList: make(map[string]stream.Viewer),
		}
	}

	h.useCase.RunIfNotRun(url)

	codecs := h.useCase.CodecGet(url)
	if codecs == nil {
		fmt.Println("    error: stream codec not found. lasterror:", h.cfg.LastError.Error())
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
		fmt.Println("    error: Muxer WriteHeader. error:", err.Error())
		return
	}

	response := stream.Response{
		Sdp64: answer,
	}

	for _, codec := range codecs {
		if codec.Type() != av.H264 &&
			codec.Type() != av.PCM_ALAW &&
			codec.Type() != av.PCM_MULAW &&
			codec.Type() != av.OPUS {
			fmt.Println("    error: track is ignored - codec not supported WebRTC. codec type:", codec.Type())
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

	go h.useCase.WritePackets(url, muxerWebRTC, audioOnly)
}
