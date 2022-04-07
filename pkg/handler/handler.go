package handler

import (
	interfaces "github.com/mikerumy/vhosting/pkg/handler/interfaces"
	methods "github.com/mikerumy/vhosting/pkg/handler/methods"
	"github.com/mikerumy/vhosting/pkg/service"
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
