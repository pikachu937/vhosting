package service

import (
	user "github.com/mikerumy/vhosting/internal/user"
)

type UserInterface interface {
	CheckUserExistence(idOrUsername interface{}) (bool, error)
	POSTUser(user user.User) error
	GETUser(id int) (*user.User, error)
	GETAllUsers() (map[int]*user.User, error)
	PATCHUser(id int, usr user.User) error
	DELETEUser(id int) error
}
