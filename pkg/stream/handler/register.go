package handler

import (
	"github.com/gin-gonic/gin"
	sconfig "github.com/mikerumy/vhosting/pkg/config_stream"
	"github.com/mikerumy/vhosting/pkg/stream"
)

func RegisterTemplateHTTPEndpoints(router *gin.Engine, cfg *sconfig.Config, uc stream.StreamUseCase) {
	h := NewStreamHandler(cfg, uc)

	router.GET("/", h.ServeIndex)
	router.GET("/stream/player/:uuid", h.ServeStreamPlayer)
}

func RegisterStreamingHTTPEndpoints(router *gin.Engine, cfg *sconfig.Config, uc stream.StreamUseCase) {
	h := NewStreamHandler(cfg, uc)

	streamRoute := router.Group("/stream")
	{
		streamRoute.GET("/codec/:uuid", h.ServeStreamCodec)
		streamRoute.POST("/receiver/:uuid", h.ServeStreamVidOverWebRTC)
		streamRoute.POST("/", h.ServeStreamWebRTC2)
	}
}
