package user

import "github.com/mikerumy/vhosting/internal/models"

type UserCommon interface {
	CreateUser(usr models.User) error
	GetUser(id int) (*models.User, error)
	GetAllUsers() (map[int]*models.User, error)
	PartiallyUpdateUser(id int, usr models.User) error
	DeleteUser(id int) error

	IsUserExists(idOrUsername interface{}) (bool, error)
}
