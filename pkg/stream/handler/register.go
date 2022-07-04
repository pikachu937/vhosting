package handler

import (
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	sconfig "github.com/dmitrij/vhosting/pkg/config_stream"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/stream"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterTemplateHTTPEndpoints(router *gin.Engine, cfg *config.Config, scfg *sconfig.Config, uc stream.StreamUseCase,
	uuc user.UserUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewStreamHandler(cfg, scfg, uc, uuc, luc, auc, suc)

	router.GET("/stream", h.ServeIndex)
	router.GET("/stream/:uuid", h.ServeStream)
}

func RegisterStreamingHTTPEndpoints(router *gin.Engine, cfg *config.Config, scfg *sconfig.Config, uc stream.StreamUseCase,
	uuc user.UserUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewStreamHandler(cfg, scfg, uc, uuc, luc, auc, suc)

	streamRoute := router.Group("/stream")
	{
		streamRoute.GET("/codec/:uuid", h.ServeStreamCodec)
		streamRoute.POST("/receiver/:uuid", h.ServeStreamVidOverWebRTC)
		streamRoute.POST("/", h.ServeStreamWebRTC2)

		streamRoute.GET("/get/:id", h.GetStream)
		streamRoute.GET("/get/all", h.GetAllStreams)
	}
}
