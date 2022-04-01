package handler

import (
	"github.com/gin-gonic/gin"
	service "github.com/mikerumy/vhservice/pkg/service/userinterface"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	editUsers := router.Group("/userinterface")
	{
		editUsers.POST("/", h.POSTUser)
		editUsers.GET("/:id", h.GETUser)
		editUsers.PUT("/:id", h.PUTUser)
		editUsers.PATCH("/:id", h.PATCHUser)
		editUsers.DELETE("/:id", h.DELETEUser)
	}

	return router
}
