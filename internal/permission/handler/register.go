package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	perm "github.com/mikerumy/vhosting/internal/permission"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc perm.PermUseCase,
	luc lg.LogUseCase, auc auth.AuthUseCase, uuc user.UserUseCase,
	suc sess.SessUseCase, ug ug.UGUseCase) {
	h := NewPermHandler(uc, luc, auc, uuc, suc, ug)

	permInterface := router.Group("/permission-interface")
	{
		permInterface.POST("", h.CreatePermission)
	}
}
