package permission

import "github.com/gin-gonic/gin"

type PermUseCase interface {
	PermCommon

	CreatePermission(ctx *gin.Context, permission Permission, timestamp string) error

	IsRequiredEmpty(id int, name, codeName string) bool
	BindJSONPermission(ctx *gin.Context) (Permission, error)
}
