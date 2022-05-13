package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
	up "github.com/mikerumy/vhosting/internal/userperm"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc up.UPUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uguc ug.UGUseCase, uuc user.UserUseCase) {
	h := NewUPHandler(uc, luc, auc, suc, uguc, uuc)

	userPermissionInterface := router.Group("/user-permission-interface")
	{
		userPermissionInterface.GET(":id", h.GetUserPermissions)
	}
}
