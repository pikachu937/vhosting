package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
)

type UserUseCase interface {
	UserCommon

	CreateUser(ctx *gin.Context, usr models.User) error
	BindJSONUser(ctx *gin.Context) (models.User, error)
	AtoiRequestedId(ctx *gin.Context) (int, error)
}
