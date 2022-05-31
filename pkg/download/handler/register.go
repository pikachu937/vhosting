package handler

import (
	"github.com/gin-gonic/gin"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/download"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc download.DownloadUseCase, luc lg.LogUseCase) {
	h := NewDownloadHandler(uc, luc)

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET("/:file_dir/:file_name", h.DownloadFile)
	}
}
