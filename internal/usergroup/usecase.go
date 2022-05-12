package usergroup

import "github.com/mikerumy/vhosting/internal/user"

type UGUseCase interface {
	UGCommon

	CreateUsergroup(usr *user.User) error
	UpdateUsergroup(usr *user.User) error
}
