package usecase

import (
	"fmt"

	perm "github.com/mikerumy/vhosting/internal/permission"
)

func (u *PermUseCase) SetGroupPermissions(id int, permIds *perm.PermIds) error {
	values := ""
	for _, val := range permIds.Ids {
		values += fmt.Sprintf("(%d,%d),", id, val)
	}
	values = values[:len(values)-1]
	return u.permRepo.SetGroupPermissions(values)
}

func (u *PermUseCase) GetGroupPermissions(id int) (*perm.PermIds, error) {
	return u.permRepo.GetGroupPermissions(id)
}

func (u *PermUseCase) DeleteGroupPermissions(id int, permIds *perm.PermIds) error {
	condIds := ""
	for _, val := range permIds.Ids {
		condIds += fmt.Sprintf("%d,", val)
	}
	condIds = condIds[:len(condIds)-1]
	return u.permRepo.DeleteGroupPermissions(id, condIds)
}
