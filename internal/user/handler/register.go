package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc user.UserUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uguc ug.UGUseCase) {
	h := NewUserHandler(uc, luc, auc, suc, uguc)

	userInterface := router.Group("/user-interface")
	{
		userInterface.POST("", h.CreateUser)
		userInterface.GET(":id", h.GetUser)
		userInterface.GET("all", h.GetAllUsers)
		userInterface.PATCH(":id", h.PartiallyUpdateUser)
		userInterface.DELETE(":id", h.DeleteUser)
	}
}
