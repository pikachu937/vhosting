package handler

import (
	"github.com/gin-gonic/gin"
	service "github.com/mikerumy/vhservice/pkg/service/userinterface"
)

type UserInterface interface {
	POSTUser(c *gin.Context)
	GETUser(c *gin.Context)
	GETAllUsers(c *gin.Context)
	PUTUser(c *gin.Context)
	PATCHUser(c *gin.Context)
	DELETEUser(c *gin.Context)
}

type Handler struct {
	UserInterface
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		UserInterface: NewUserInterfaceHandler(services),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	userInterface := router.Group("/userinterface")
	{
		userInterface.POST("/", h.POSTUser)
		userInterface.GET("/:id", h.GETUser)
		userInterface.GET("/all", h.GETAllUsers)
		userInterface.PUT("/:id", h.PUTUser)
		userInterface.PATCH("/:id", h.PATCHUser)
		userInterface.DELETE("/:id", h.DELETEUser)
	}

	return router
}
