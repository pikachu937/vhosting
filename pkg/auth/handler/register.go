package handler

import (
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc auth.AuthUseCase, uuc user.UserUseCase,
	suc sess.SessUseCase, luc logger.LogUseCase) {
	h := NewAuthHandler(cfg, uc, uuc, suc, luc)

	authRoute := router.Group("/auth")
	{
		authRoute.POST("/signin", h.SignIn)
		authRoute.POST("/change_password", h.ChangePassword)
		authRoute.GET("/signout", h.SignOut)
	}
}
