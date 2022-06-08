package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/pkg/download"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc download.DownloadUseCase, luc logger.LogUseCase) {
	h := NewDownloadHandler(uc, luc)

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET("/:file_dir/:file_name", h.DownloadFile)
	}
}
