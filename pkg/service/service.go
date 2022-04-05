package service

import (
	"github.com/mikerumy/vhservice/pkg/repository"
	service "github.com/mikerumy/vhservice/pkg/service/userinterface"
)

type Service struct {
	service.UserInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserInterface: service.NewUserInterfaceService(repos.UserInterface),
	}
}
