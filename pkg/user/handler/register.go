package handler

import (
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc user.UserUseCase,
	luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewUserHandler(cfg, uc, luc, auc, suc)

	userRoute := router.Group("/user")
	{
		userRoute.POST("", h.CreateUser)
		userRoute.GET(":id", h.GetUser)
		userRoute.GET("all", h.GetAllUsers)
		userRoute.POST("/change_password", h.UpdateUserPassword)
		userRoute.PATCH(":id", h.PartiallyUpdateUser)
		userRoute.DELETE(":id", h.DeleteUser)
	}
}
