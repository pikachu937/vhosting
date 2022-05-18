package permission

import (
	"github.com/gin-gonic/gin"
)

type PermCommon interface {
	GetAllPermissions() (map[int]*Perm, error)
	GetUserPermissions(id int) (*PermIds, error)
	GetGroupPermissions(id int) (*PermIds, error)
}

type PermUseCase interface {
	PermCommon

	SetUserPermissions(id int, perms PermIds) error
	DeleteUserPermissions(id int, perms PermIds) error

	SetGroupPermissions(id int, perms PermIds) error
	DeleteGroupPermissions(id int, perms PermIds) error

	BindJSONPerms(ctx *gin.Context) (PermIds, error)
	IsRequiredEmpty(perms PermIds) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type PermRepository interface {
	PermCommon

	SetUserPermissions(values string) error
	DeleteUserPermissions(id int, condIds string) error

	SetGroupPermissions(values string) error
	DeleteGroupPermissions(id int, condIds string) error
}
