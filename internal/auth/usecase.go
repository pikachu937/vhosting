package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthUseCase interface {
	AuthCommon

	ReadCookie(ctx *gin.Context) string
	BindJSONNamepass(ctx *gin.Context) (Namepass, error)
	GenerateToken(namepass Namepass) (string, error)
	SendCookie(ctx *gin.Context, token string)
	ParseToken(token string) (Namepass, error)
	DeleteCookie(ctx *gin.Context)
}
