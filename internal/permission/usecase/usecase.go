package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	perm "github.com/mikerumy/vhosting/internal/permission"
	"github.com/mikerumy/vhosting/pkg/config"
)

type PermUseCase struct {
	cfg      config.Config
	permRepo perm.PermRepository
}

func NewPermUseCase(cfg config.Config, permRepo perm.PermRepository) *PermUseCase {
	return &PermUseCase{
		cfg:      cfg,
		permRepo: permRepo,
	}
}

func (u *PermUseCase) GetAllPermissions() (map[int]*perm.Perm, error) {
	return u.permRepo.GetAllPermissions()
}

func (u *PermUseCase) BindJSONPermIds(ctx *gin.Context) (perm.PermIds, error) {
	var permIds perm.PermIds
	if err := ctx.BindJSON(&permIds); err != nil {
		return permIds, err
	}
	return permIds, nil
}

func (u *PermUseCase) IsRequiredEmpty(perms perm.PermIds) bool {
	if len(perms.Ids) == 0 {
		return true
	}
	return false
}

func (u *PermUseCase) AtoiRequestedId(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return -1, err
	}
	return id, nil
}
