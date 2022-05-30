package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/pkg/download"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type DownloadHandler struct{}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}

func (h *DownloadHandler) DownloadFile(ctx *gin.Context) {
	log := logger.Setup(ctx)

	fmt.Println(ctx.ClientIP())

	var filename download.Download
	filename.Url = ctx.Param("file_name")

	extension := filename.Url[len(filename.Url)-4:]
	strings.ToLower(extension)
	if extension != ".mp4" {
		log.Message = "extension not mp4"
		fmt.Println(log.Message)
		responder.Response(ctx, log)
		return
	}
	ctx.File("./files/" + filename.Url)

}
