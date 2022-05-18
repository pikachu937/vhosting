package group

import "github.com/gin-gonic/gin"

type GroupCommon interface {
	CreateGroup(grp Group) error
	GetGroup(id int) (*Group, error)
	GetAllGroups() (map[int]*Group, error)
	PartiallyUpdateGroup(grp *Group) error
	DeleteGroup(id int) error

	IsGroupExists(idOrName interface{}) (bool, error)
}

type GroupUseCase interface {
	GroupCommon

	BindJSONGroup(ctx *gin.Context) (Group, error)
	IsRequiredEmpty(name string) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type GroupRepository interface {
	GroupCommon
}
