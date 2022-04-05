package service

import (
	interfaces "github.com/mikerumy/vhservice/pkg/service/interfaces"
	methods "github.com/mikerumy/vhservice/pkg/service/methods"
	"github.com/mikerumy/vhservice/pkg/storage"
)

type Service struct {
	interfaces.UserInterface
	interfaces.Authorization
}

func NewService(stor *storage.Storage) *Service {
	return &Service{
		UserInterface: methods.NewUserInterfaceService(stor.UserInterface),
		Authorization: methods.NewAuthService(stor.Authorization),
	}
}
