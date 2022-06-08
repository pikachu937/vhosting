package handler

import (
	"github.com/gin-gonic/gin"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/download"
	"github.com/mikerumy/vhosting/pkg/logger"
)

type DownloadHandler struct {
	useCase    download.DownloadUseCase
	logUseCase logger.LogUseCase
}

func NewDownloadHandler(useCase download.DownloadUseCase,
	logUseCase logger.LogUseCase) *DownloadHandler {
	return &DownloadHandler{
		useCase:    useCase,
		logUseCase: logUseCase,
	}
}

func (h *DownloadHandler) DownloadFile(ctx *gin.Context) {
	log := logger.Init(ctx)

	fileName := ctx.Param("file_name")

	if !h.useCase.IsValidExtension(fileName) {
		h.logUseCase.Report(ctx, log, msg.ErrorExtensionIsNotMp4())
	}

	fileDir := ctx.Param("file_dir")

	download := h.useCase.CreateDownloadLink(fileDir + "/" + fileName)

	h.logUseCase.Report(ctx, log, msg.InfoPutDownloadLink(download))
}
