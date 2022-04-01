package service

import (
	vhs "github.com/mikerumy/vhservice"
	"github.com/mikerumy/vhservice/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserEdit interface {
	CreateUser(user vhs.User) (int, error)
	GetUser(id int) (*vhs.User, error)
	UpdateUser(id int, user vhs.User) (int, error)
	PartiallyUpdateUser(id int, user vhs.User) (int, error)
	DeleteUser(id int) (int, error)
}

type Service struct {
	UserEdit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserEdit: NewUserEditService(repos.UserEdit),
	}
}
