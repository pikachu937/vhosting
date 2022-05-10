package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc user.UserUseCase, luc logging.LoggingUseCase) {
	h := NewUserHandler(uc, luc)

	userInterface := router.Group("/user-interface")
	{
		userInterface.POST("", h.CreateUser)
		userInterface.GET(":id", h.GetUser)
		userInterface.GET("all", h.GetAllUsers)
		userInterface.PATCH(":id", h.PartiallyUpdateUser)
		userInterface.DELETE(":id", h.DeleteUser)
	}
}
