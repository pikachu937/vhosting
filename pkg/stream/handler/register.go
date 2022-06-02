package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/stream"
)

func RegisterTemplateHTTPEndpoints(router *gin.Engine, uc stream.StreamUseCase, cfg *models.ConfigST) {
	h := NewStreamHandler(uc, cfg)

	router.GET("/", h.ServeIndex)
	router.GET("/stream/player/:uuid", h.ServeStreamPlayer)
}

func RegisterStreamingHTTPEndpoints(router *gin.Engine, uc stream.StreamUseCase, cfg *models.ConfigST) {
	h := NewStreamHandler(uc, cfg)

	streamRoute := router.Group("/stream")
	{
		streamRoute.GET("/codec/:uuid", h.ServeStreamCodec)
		streamRoute.POST("/receiver/:uuid", h.ServeStreamWebRTC)
		streamRoute.POST("/", h.ServeStreamWebRTC2)
	}
}
