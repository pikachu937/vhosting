package permission

import "github.com/gin-gonic/gin"

type PermUseCase interface {
	PermCommon

	IsRequiredEmpty(id int, name, codeName string) bool
	BindJSONPermission(ctx *gin.Context) (Permission, error)
}
