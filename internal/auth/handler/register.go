package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.AuthUseCase, uuc user.UserUseCase, suc session.SessionUseCase, luc logging.LoggingUseCase) {
	h := NewAuthHandler(uc, uuc, suc, luc)

	authorization := router.Group("/auth")
	{
		authorization.POST("/sign-in", h.SignIn)
		authorization.POST("/change-password", h.ChangePassword)
		authorization.POST("/sign-out", h.SignOut)
	}
}
