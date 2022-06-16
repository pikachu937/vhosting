package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/deepch/vdk/av"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
	"github.com/gin-gonic/gin"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/config"
	sconfig "github.com/mikerumy/vhosting/pkg/config_stream"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/stream"
	"github.com/mikerumy/vhosting/pkg/timedate"
	"github.com/mikerumy/vhosting/pkg/user"
)

type StreamHandler struct {
	cfg         *config.Config
	scfg        *sconfig.Config
	useCase     stream.StreamUseCase
	userUseCase user.UserUseCase
	logUseCase  logger.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
}

func NewStreamHandler(cfg *config.Config, scfg *sconfig.Config, useCase stream.StreamUseCase,
	userUseCase user.UserUseCase, logUseCase logger.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase) *StreamHandler {
	return &StreamHandler{
		cfg:         cfg,
		scfg:        scfg,
		useCase:     useCase,
		userUseCase: userUseCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
	}
}

func (h *StreamHandler) GetStream(ctx *gin.Context) {
	actPermission := "get_stream"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check user existence, get user
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsStreamExists(reqId)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotCheckStreamExistence(err))
		return
	}
	if !exists {
		h.logUseCase.Report(ctx, log, msg.ErrorStreamWithRequestedIDIsNotExist())
		return
	}

	gottenStream, err := h.useCase.GetStream(reqId)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotGetStream(err))
		return
	}

	h.logUseCase.Report(ctx, log, msg.InfoGotStream(gottenStream))
}

func (h *StreamHandler) GetAllStreams(ctx *gin.Context) {
	actPermission := "get_all_streams"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	urlparams := h.useCase.ParseURLParams(ctx)

	// Get all users. If gotten is nothing - send such a message
	gottenStreams, err := h.useCase.GetAllStreams(urlparams)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotGetAllStreams(err))
		return
	}

	if gottenStreams == nil {
		h.logUseCase.Report(ctx, log, msg.InfoNoStreamsAvailable())
		return
	}

	h.logUseCase.Report(ctx, log, msg.InfoGotAllStreams(gottenStreams))
}

func (h *StreamHandler) ServeIndex(ctx *gin.Context) {
	_, list := h.useCase.List()
	if len(list) > 0 {
		ctx.Header("Cache-Control", "no-cache, max-age=0, must-revalidate, no-store")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Redirect(http.StatusMovedPermanently, "stream/player/"+list[0])
	} else {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"port":    h.scfg.Server.HTTPPort,
			"version": time.Now().String(),
		})
	}
}

func (h *StreamHandler) ServeStreamPlayer(ctx *gin.Context) {
	_, list := h.useCase.List()
	sort.Strings(list)
	ctx.HTML(http.StatusOK, "player.tmpl", gin.H{
		"port":     h.scfg.Server.HTTPPort,
		"suuid":    ctx.Param("uuid"),
		"suuidMap": list,
		"version":  time.Now().String(),
	})
}

func (h *StreamHandler) ServeStreamCodec(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if !h.useCase.Exit(uuid) {
		return
	}

	h.useCase.RunIfNotRun(uuid)

	codecs := h.useCase.CodecGet(uuid)
	if codecs == nil {
		return
	}

	var tmpCodec []stream.JCodec
	for _, codec := range codecs {
		if codec.Type() != av.H264 && codec.Type() != av.PCM_ALAW && codec.Type() != av.PCM_MULAW && codec.Type() != av.OPUS {
			logger.Printc(ctx, msg.ErrorTrackIsIgnoredCodecNotSupportedWebRTC(codec.Type()))
			continue
		}

		if codec.Type().IsVideo() {
			tmpCodec = append(tmpCodec, stream.JCodec{Type: "video"})
		} else {
			tmpCodec = append(tmpCodec, stream.JCodec{Type: "audio"})
		}
	}

	b, err := json.Marshal(tmpCodec)
	if err != nil {
		return
	}

	_, err = ctx.Writer.Write(b)
	if err != nil {
		logger.Printc(ctx, msg.ErrorWritingOfCodecError(err))
	}
}

func (h *StreamHandler) ServeStreamVidOverWebRTC(ctx *gin.Context) {
	suuid := ctx.PostForm("suuid")
	if !h.useCase.Exit(suuid) {
		logger.Printc(ctx, msg.InfoStreamNotFound(suuid))
		return
	}

	h.useCase.RunIfNotRun(suuid)

	codecs := h.useCase.CodecGet(suuid)
	if codecs == nil {
		logger.Printc(ctx, msg.InfoStreamCodecNotFound(suuid))
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
		logger.Printc(ctx, msg.ErrorWriteHeaderError(err))
		return
	}

	if _, err := ctx.Writer.Write([]byte(answer)); err != nil {
		logger.Printc(ctx, msg.ErrorCannotWriteBytes(err))
		return
	}

	go h.useCase.WritePackets(suuid, muxerWebRTC, audioOnly)
}

func (h *StreamHandler) ServeStreamWebRTC2(ctx *gin.Context) {
	url := ctx.PostForm("url")
	if _, ok := h.scfg.Streams[url]; !ok {
		h.scfg.Streams[url] = stream.StreamSettings{
			URL:        url,
			OnDemand:   true,
			ClientList: make(map[string]stream.Viewer),
		}
	}

	h.useCase.RunIfNotRun(url)

	codecs := h.useCase.CodecGet(url)
	if codecs == nil {
		logger.Printc(ctx, msg.ErrorStreamCodecNotFound(h.scfg.LastError))
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
		logger.Printc(ctx, msg.ErrorMuxerWriteHeaderError(err))
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
			logger.Printc(ctx, msg.ErrorTrackIsIgnoredCodecNotSupportedWebRTC(codec.Type()))
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

func (h *StreamHandler) isPermsGranted_getUserId(ctx *gin.Context, log *logger.Log, permission string) (bool, int) {
	headerToken := h.authUseCase.ReadHeader(ctx)
	if !h.authUseCase.IsTokenExists(headerToken) {
		h.logUseCase.Report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	session, err := h.sessUseCase.GetSessionAndDate(headerToken)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotGetSessionAndDate(err))
		return false, -1
	}
	if !h.authUseCase.IsSessionExists(session) {
		h.logUseCase.Report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	if timedate.IsDateExpired(session.CreationDate, h.cfg.SessionTTLHours) {
		if err := h.sessUseCase.DeleteSession(headerToken); err != nil {
			h.logUseCase.Report(ctx, log, msg.ErrorCannotDeleteSession(err))
			return false, -1
		}
		h.logUseCase.Report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	headerNamepass, err := h.authUseCase.ParseToken(headerToken)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotParseToken(err))
		return false, -1
	}

	gottenUserId, err := h.userUseCase.GetUserId(headerNamepass.Username)
	if err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return false, -1
	}
	if gottenUserId < 0 {
		if err := h.sessUseCase.DeleteSession(headerToken); err != nil {
			h.logUseCase.Report(ctx, log, msg.ErrorCannotDeleteSession(err))
			return false, -1
		}
		h.logUseCase.Report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false, -1
	}

	log.SessionOwner = headerNamepass.Username

	isSUorStaff := false
	hasPersonalPerm := false
	if isSUorStaff, err = h.userUseCase.IsUserSuperuserOrStaff(headerNamepass.Username); err != nil {
		h.logUseCase.Report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false, -1
	}
	if !isSUorStaff {
		if hasPersonalPerm, err = h.userUseCase.IsUserHavePersonalPermission(gottenUserId, permission); err != nil {
			h.logUseCase.Report(ctx, log, msg.ErrorCannotCheckPersonalPermission(err))
			return false, -1
		}
	}

	if !isSUorStaff && !hasPersonalPerm {
		h.logUseCase.Report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	return true, gottenUserId
}
