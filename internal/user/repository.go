package user

import "github.com/mikerumy/vhosting/internal/models"

type UserRepository interface {
	UserCommon

	CreateUser(usr models.User) error
}
