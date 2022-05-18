package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc group.GroupUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewGroupHandler(uc, luc, auc, suc, uuc)

	groupInterface := router.Group("/group-interface")
	{
		groupInterface.POST("", h.CreateGroup)
		groupInterface.GET(":id", h.GetGroup)
		groupInterface.GET("all", h.GetAllGroups)
		groupInterface.PATCH(":id", h.PartiallyUpdateGroup)
		groupInterface.DELETE(":id", h.DeleteGroup)
	}
}
