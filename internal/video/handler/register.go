package handler

import (
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/internal/video"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc video.VideoUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewVideoHandler(cfg, uc, luc, auc, suc, uuc)

	videoRoute := router.Group("/video")
	{
		videoRoute.POST("", h.CreateVideo)
		videoRoute.GET(":id", h.GetVideo)
		videoRoute.GET("all", h.GetAllVideos)
		videoRoute.PATCH(":id", h.PartiallyUpdateVideo)
		videoRoute.DELETE(":id", h.DeleteVideo)

		videoRoute.GET("/concat", h.VideoConcat)
	}

	router.GET("/t", h.Test)
}
