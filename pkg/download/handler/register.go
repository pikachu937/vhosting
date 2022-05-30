package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine) {
	h := NewDownloadHandler()

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET(":file_name", h.DownloadFile)
	}
}
