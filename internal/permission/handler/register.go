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

	permRoute := router.Group("/perm")
	{
		permRoute.GET("all", h.GetAllPermissions)
	}

	permSetUserRoute := router.Group("/perm/user")
	{
		permSetUserRoute.POST(":id", h.SetUserPermissions)
		permSetUserRoute.GET(":id", h.GetUserPermissions)
		permSetUserRoute.DELETE(":id", h.DeleteUserPermissions)
	}

	permSetGroupRoute := router.Group("/perm/group")
	{
		permSetGroupRoute.POST(":id", h.SetGroupPermissions)
		permSetGroupRoute.GET(":id", h.GetGroupPermissions)
		permSetGroupRoute.DELETE(":id", h.DeleteGroupPermissions)
	}
}