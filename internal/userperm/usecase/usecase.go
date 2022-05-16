package usecase

import (
	up "github.com/mikerumy/vhosting/internal/userperm"
)

type UPUseCase struct {
	upRepo up.UPRepository
}

func NewUPUseCase(upRepo up.UPRepository) *UPUseCase {
	return &UPUseCase{
		upRepo: upRepo,
	}
}

func (u *UPUseCase) CreateUserperm(userperm *up.Userperm) error {
	return u.upRepo.CreateUserperm(userperm)
}

func (u *UPUseCase) GetUserPermissions(id int) (map[int]*up.Userperm, error) {
	return u.upRepo.GetUserPermissions(id)
}

func (u *UPUseCase) UpsertUserPermissions(userId, groupId int) error {
	return u.upRepo.UpsertUserPermissions(userId, groupId)
}

func (u *UPUseCase) DeleteUserPermissions(userId, groupId int) error {
	return u.upRepo.DeleteUserPermissions(userId, groupId)
}
