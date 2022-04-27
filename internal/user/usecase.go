package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting2/internal/models"
)

type UserUseCase interface {
	UserCommon

	BindJSONUser(ctx *gin.Context) (models.User, error)
	AtoiRequestedId(ctx *gin.Context) (int, error)
}
