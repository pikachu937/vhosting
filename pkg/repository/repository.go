package repository

import (
	vhs "github.com/mikerumy/vhservice"
	repository "github.com/mikerumy/vhservice/pkg/repository/userinterface"
)

type Repository struct {
	repository.UserInterface
}

func NewRepository(cfg vhs.DBConfig) *Repository {
	return &Repository{
		UserInterface: repository.NewUserInterfaceRepo(cfg),
	}
}
