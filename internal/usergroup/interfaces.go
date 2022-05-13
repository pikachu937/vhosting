package usergroup

import "github.com/mikerumy/vhosting/internal/user"

// type UGCommon interface {
// 	IsUserInGroup(userId, groupId int) bool
// }

type UGUseCase interface {
	// UGCommon

	CreateUsergroup(usr *user.User) error
	UpdateUsergroup(usr *user.User) error
}

type UGRepository interface {
	// UGCommon

	CreateUsergroup(userId, groupId int) error
	UpdateUsergroup(userId, groupId int) error
}
