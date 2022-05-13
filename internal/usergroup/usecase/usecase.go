package usecase

import (
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
)

type UGUseCase struct {
	ugRepo ug.UGRepository
}

func NewUGUseCase(ugRepo ug.UGRepository) *UGUseCase {
	return &UGUseCase{
		ugRepo: ugRepo,
	}
}

// func (u *UGUseCase) IsUserInGroup(userId, groupId int) bool {
// 	return u.ugRepo.IsUserInGroup(userId, groupId)
// }

func (u *UGUseCase) CreateUsergroup(usr *user.User) error {
	if usr.IsStaff {
		return u.ugRepo.CreateUsergroup(usr.Id, ug.StaffGroup)
	}
	return u.ugRepo.CreateUsergroup(usr.Id, ug.UserGroup)
}

func (u *UGUseCase) UpdateUsergroup(usr *user.User) error {
	if usr.IsStaff {
		return u.ugRepo.UpdateUsergroup(usr.Id, ug.StaffGroup)
	}
	return u.ugRepo.UpdateUsergroup(usr.Id, ug.UserGroup)
}
