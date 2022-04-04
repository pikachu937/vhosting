package repository

import (
	vhs "github.com/mikerumy/vhservice"
)

type UserInterface interface {
	POSTUser(user vhs.User) (int, error)
	GETUser(id int) (*vhs.User, error)
	PUTUser(id int, user vhs.User) (int, error)
	PATCHUser(id int, user vhs.User) (int, error)
	DELETEUser(id int) (int, error)
}

type Repository struct {
	UserInterface
}

func NewRepository(cfg vhs.DBConfig) *Repository {
	return &Repository{
		UserInterface: NewUserInterfaceRepo(cfg),
	}
}
