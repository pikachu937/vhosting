package usecase

import (
	"github.com/gin-gonic/gin"
	perm "github.com/mikerumy/vhosting/internal/permission"
)

type PermUseCase struct {
	permRepo perm.PermRepository
}

func NewPermUseCase(permRepo perm.PermRepository) *PermUseCase {
	return &PermUseCase{
		permRepo: permRepo,
	}
}

func (u *PermUseCase) CreatePermission(permission perm.Permission) error {
	return u.permRepo.CreatePermission(permission)
}

func (u *PermUseCase) IsEmpty(id int, name, codeName string) bool {
	if id == 0 || name == "" || codeName == "" {
		return true
	}
	return false
}

func (u *PermUseCase) BindJSONPermission(ctx *gin.Context) (perm.Permission, error) {
	var permission perm.Permission
	if err := ctx.BindJSON(&permission); err != nil {
		return permission, err
	}
	return permission, nil
}
