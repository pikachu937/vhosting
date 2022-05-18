package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/group"
	"github.com/mikerumy/vhosting/pkg/config_tool"
)

type GroupUseCase struct {
	cfg       config_tool.Config
	groupRepo group.GroupRepository
}

func NewGroupUseCase(cfg config_tool.Config, groupRepo group.GroupRepository) *GroupUseCase {
	return &GroupUseCase{
		cfg:       cfg,
		groupRepo: groupRepo,
	}
}

func (u *GroupUseCase) CreateGroup(grp group.Group) error {
	return u.groupRepo.CreateGroup(grp)
}

func (u *GroupUseCase) GetGroup(id int) (*group.Group, error) {
	return u.groupRepo.GetGroup(id)
}

func (u *GroupUseCase) GetAllGroups() (map[int]*group.Group, error) {
	return u.groupRepo.GetAllGroups()
}

func (u *GroupUseCase) PartiallyUpdateGroup(grp *group.Group) error {
	return u.groupRepo.PartiallyUpdateGroup(grp)
}

func (u *GroupUseCase) DeleteGroup(id int) error {
	return u.groupRepo.DeleteGroup(id)
}

func (u *GroupUseCase) IsGroupExists(idOrName interface{}) (bool, error) {
	var err error
	exists, err := u.groupRepo.IsGroupExists(idOrName)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *GroupUseCase) BindJSONGroup(ctx *gin.Context) (group.Group, error) {
	var err error
	var grp group.Group
	if err = ctx.BindJSON(&grp); err != nil {
		return grp, err
	}
	return grp, nil
}

func (u *GroupUseCase) IsRequiredEmpty(name string) bool {
	if name == "" {
		return true
	}
	return false
}

func (u *GroupUseCase) AtoiRequestedId(ctx *gin.Context) (int, error) {
	var err error
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return -1, err
	}
	return id, nil
}
