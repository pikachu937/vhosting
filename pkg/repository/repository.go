package repository

import (
	vhs "github.com/mikerumy/vhservice"
	interfaces "github.com/mikerumy/vhservice/pkg/repository/interfaces"
	methods "github.com/mikerumy/vhservice/pkg/repository/methods"
)

type Repository struct {
	interfaces.UserInterface
	interfaces.Authorization
}

func NewRepository(cfg vhs.DBConfig) *Repository {
	return &Repository{
		UserInterface: methods.NewUserInterfaceRepo(cfg),
		Authorization: methods.NewAuthorizationRepo(cfg),
	}
}
