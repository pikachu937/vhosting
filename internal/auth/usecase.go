package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
)

type AuthUseCase interface {
	AuthCommon

	CreateSession(ctx *gin.Context, username, timestamp string) error
	ReadCookie(ctx *gin.Context) string
	BindJSONNamepass(ctx *gin.Context) (models.Namepass, error)
	GenerateToken(namepass models.Namepass) (string, error)
	SendCookie(ctx *gin.Context, token string)
	ParseToken(token string) (models.Namepass, error)
	DeleteCookie(ctx *gin.Context)
}
