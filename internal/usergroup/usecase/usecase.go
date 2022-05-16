package usecase

import (
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

func (u *UGUseCase) CreateUsergroup(userId, groupId int) error {
	return u.ugRepo.CreateUsergroup(userId, groupId)
}

func (u *UGUseCase) DeleteUsergroup(userId, groupId int) error {
	return u.ugRepo.DeleteUsergroup(userId, groupId)
}

// func (u *UGUseCase) IsUserInGroup(userId, groupId int) bool {
// 	return u.ugRepo.IsUserInGroup(userId, groupId)
// }

// func (u *UGUseCase) UpdateUsergroup(usr *user.User) error {
// 	if usr.IsStaff {
// 		return u.ugRepo.UpdateUsergroup(usr.Id, ug.StaffGroup)
// 	}
// 	return u.ugRepo.UpdateUsergroup(usr.Id, ug.UserGroup)
// }
