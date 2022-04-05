package handler

import (
	interfaces "github.com/mikerumy/vhservice/pkg/handler/interfaces"
	methods "github.com/mikerumy/vhservice/pkg/handler/methods"
	"github.com/mikerumy/vhservice/pkg/service"
)

type Handler struct {
	interfaces.UserInterface
	interfaces.Authorization
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		UserInterface: methods.NewUserInterfaceHandler(services),
		Authorization: methods.NewAuthorizationHandler(services),
	}
}
