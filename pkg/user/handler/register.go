package handler

import (
	"github.com/gin-gonic/gin"
	lg "github.com/mikerumy/vhosting/internal/logging"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc user.UserUseCase, luc lg.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewUserHandler(uc, luc, auc, suc)

	userRoute := router.Group("/user")
	{
		userRoute.POST("", h.CreateUser)
		userRoute.GET(":id", h.GetUser)
		userRoute.GET("all", h.GetAllUsers)
		// userRoute.POST(":id", h.UpdateUserPassword)
		userRoute.PATCH(":id", h.PartiallyUpdateUser)
		userRoute.DELETE(":id", h.DeleteUser)
	}
}
