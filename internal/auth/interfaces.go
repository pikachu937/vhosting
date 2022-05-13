package auth

import "github.com/gin-gonic/gin"

type AuthCommon interface {
	GetNamepass(namepass Namepass) error
	UpdateNamepassPassword(namepass Namepass) error
	IsNamepassExists(username, passwordHash string) (bool, error)
}

type AuthUseCase interface {
	AuthCommon

	ReadCookie(ctx *gin.Context) string
	IsTokenExist(token string) bool
	IsMatched(username_1, username_2 string) bool
	BindJSONNamepass(ctx *gin.Context) (Namepass, error)
	GenerateToken(namepass Namepass) (string, error)
	SendCookie(ctx *gin.Context, token string)
	ParseToken(token string) (Namepass, error)
	DeleteCookie(ctx *gin.Context)
}

type AuthRepository interface {
	AuthCommon

	UpdateNamepassLastLogin(username, token string) error
}
