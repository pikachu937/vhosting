package user

import (
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	UserCommon

	IsRequiredEmpty(username, password string) bool

	BindJSONUser(ctx *gin.Context) (User, error)
	AtoiRequestedId(ctx *gin.Context) (int, error)
}
