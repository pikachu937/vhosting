package handler

import (
	"github.com/gin-gonic/gin"
	lg "github.com/mikerumy/vhosting/internal/logging"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/video"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc video.VideoUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewVideoHandler(uc, luc, auc, suc, uuc)

	videoRoute := router.Group("/video")
	{
		videoRoute.POST("", h.CreateVideo)
		videoRoute.GET(":id", h.GetVideo)
		videoRoute.GET("all", h.GetAllVideos)
		videoRoute.PATCH(":id", h.PartiallyUpdateVideo)
		videoRoute.DELETE(":id", h.DeleteVideo)
	}
}
