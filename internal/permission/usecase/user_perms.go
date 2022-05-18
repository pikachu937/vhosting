package usecase

import (
	"fmt"

	perm "github.com/mikerumy/vhosting/internal/permission"
)

func (u *PermUseCase) SetUserPermissions(id int, perms perm.PermIds) error {
	values := ""
	for _, val := range perms.Ids {
		values += fmt.Sprintf("(%d,%d),", id, val)
	}
	values = values[:len(values)-1]
	return u.permRepo.SetUserPermissions(values)
}

func (u *PermUseCase) GetUserPermissions(id int) (*perm.PermIds, error) {
	return u.permRepo.GetUserPermissions(id)
}

func (u *PermUseCase) DeleteUserPermissions(id int, perms perm.PermIds) error {
	condIds := ""
	for _, val := range perms.Ids {
		condIds += fmt.Sprintf("%d,", val)
	}
	condIds = condIds[:len(condIds)-1]

	return u.permRepo.DeleteUserPermissions(id, condIds)
}
