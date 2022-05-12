package session

import "github.com/gin-gonic/gin"

type SessUseCase interface {
	SessCommon

	CreateSession(ctx *gin.Context, username, token, timestamp string) error
}
