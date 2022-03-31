package repository

import (
	"github.com/jmoiron/sqlx"
	vhs "github.com/mikerumy/vhservice"
)

type UserEdit interface {
	CreateUser(user vhs.User) (int, error)
	GetUser(id int) (*vhs.User, error)
	UpdateUser(id int, user vhs.User) (int, error)
	DeleteUser(id int) (int, error)
}

type Repository struct {
	UserEdit
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserEdit: NewUserEditPostgres(db),
	}
}
