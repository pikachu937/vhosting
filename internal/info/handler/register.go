package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/info"
	lg "github.com/mikerumy/vhosting/internal/logging"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc info.InfoUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewInfoHandler(uc, luc, auc, suc, uuc)

	infoRoute := router.Group("/info")
	{
		infoRoute.POST("", h.CreateInfo)
		infoRoute.GET(":id", h.GetInfo)
		infoRoute.GET("all", h.GetAllInfos)
		infoRoute.PATCH(":id", h.PartiallyUpdateInfo)
		infoRoute.DELETE(":id", h.DeleteInfo)
	}
}
