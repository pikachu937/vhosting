package handler

import (
	handler "github.com/mikerumy/vhservice/pkg/handler/userinterface"
	"github.com/mikerumy/vhservice/pkg/service"
)

type Handler struct {
	handler.UserInterface
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		UserInterface: handler.NewUserInterfaceHandler(services),
	}
}
