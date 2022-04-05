package service

import (
	vhs "github.com/mikerumy/vhservice"
	repository "github.com/mikerumy/vhservice/pkg/repository/userinterface"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserInterface interface {
	POSTUser(user vhs.User) (int, error)
	GETUser(id int) (*vhs.User, error)
	GETAllUsers() (map[int]*vhs.User, error)
	PUTUser(id int, user vhs.User) (int, error)
	PATCHUser(id int, user vhs.User) (int, error)
	DELETEUser(id int) (int, error)
}

type Service struct {
	UserInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserInterface: NewUserInterfaceService(repos.UserInterface),
	}
}
