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
