package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	perm "github.com/mikerumy/vhosting/internal/permission"
	"github.com/mikerumy/vhosting/pkg/config_tool"
)

type PermUseCase struct {
	cfg      config_tool.Config
	permRepo perm.PermRepository
}

func NewPermUseCase(cfg config_tool.Config, permRepo perm.PermRepository) *PermUseCase {
	return &PermUseCase{
		cfg:      cfg,
		permRepo: permRepo,
	}
}

func (u *PermUseCase) GetAllPermissions() (map[int]*perm.Perm, error) {
	return u.permRepo.GetAllPermissions()
}

func (u *PermUseCase) BindJSONPerms(ctx *gin.Context) (perm.PermIds, error) {
	var err error
	var perms perm.PermIds
	if err = ctx.BindJSON(&perms); err != nil {
		return perms, err
	}
	return perms, nil
}

func (u *PermUseCase) IsRequiredEmpty(perms perm.PermIds) bool {
	if len(perms.Ids) == 0 {
		return true
	}
	return false
}

func (u *PermUseCase) AtoiRequestedId(ctx *gin.Context) (int, error) {
	var err error
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return -1, err
	}
	return id, nil
}
