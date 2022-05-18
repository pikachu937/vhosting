package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	perm "github.com/mikerumy/vhosting/internal/permission"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc perm.PermUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase, guc group.GroupUseCase) {
	h := NewPermHandler(uc, luc, auc, suc, uuc, guc)

	permInterface := router.Group("/perm-interface")
	{
		permInterface.GET("perms", h.GetAllPermissions)

		permInterface.POST("user/:id", h.SetUserPermissions)
		permInterface.GET("user/:id", h.GetUserPermissions)
		permInterface.DELETE("user/:id", h.DeleteUserPermissions)

		permInterface.POST("group/:id", h.SetGroupPermissions)
		permInterface.GET("group/:id", h.GetGroupPermissions)
		permInterface.DELETE("group/:id", h.DeleteGroupPermissions)
	}
}
