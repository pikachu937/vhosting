package user

import (
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	UserCommon

	IsEmpty(username, password string) bool
	CreateUser(ctx *gin.Context, usr User, timestamp string) error
	BindJSONUser(ctx *gin.Context) (User, error)
	AtoiRequestedId(ctx *gin.Context) (int, error)
}
