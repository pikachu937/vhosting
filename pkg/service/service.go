package service

import (
	"github.com/mikerumy/vhservice/pkg/repository"
	interfaces "github.com/mikerumy/vhservice/pkg/service/interfaces"
	methods "github.com/mikerumy/vhservice/pkg/service/methods"
)

type Service struct {
	interfaces.UserInterface
	interfaces.Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserInterface: methods.NewUserInterfaceService(repos.UserInterface),
		Authorization: methods.NewAuthService(repos.Authorization),
	}
}
