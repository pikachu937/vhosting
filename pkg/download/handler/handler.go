package handler

import (
	"github.com/gin-gonic/gin"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/download"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type DownloadHandler struct {
	useCase    download.DownloadUseCase
	logUseCase lg.LogUseCase
}

func NewDownloadHandler(useCase download.DownloadUseCase,
	logUseCase lg.LogUseCase) *DownloadHandler {
	return &DownloadHandler{
		useCase:    useCase,
		logUseCase: logUseCase,
	}
}

func (h *DownloadHandler) DownloadFile(ctx *gin.Context) {
	log := logger.Init(ctx)

	fileName := ctx.Param("file_name")

	if !h.useCase.IsValidExtension(fileName) {
		h.report(ctx, log, msg.ErrorExtensionNotMp4())
	}

	fileDir := ctx.Param("file_dir")

	download := h.useCase.CreateDownloadLink(fileDir + "/" + fileName)

	h.report(ctx, log, msg.InfoPutDownloadLink(download))
}

func (h *DownloadHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}
