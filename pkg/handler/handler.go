package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhservice/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	editUsers := router.Group("/user-edit")
	{
		editUsers.POST("/", h.CreateUser)
		editUsers.GET("/:id", h.GetUser)
		editUsers.PUT("/:id", h.UpdateUser)
		editUsers.PATCH("/:id", h.PartiallyUpdateUser)
		editUsers.DELETE("/:id", h.DeleteUser)
	}

	return router
}
