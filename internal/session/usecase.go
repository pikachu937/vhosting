package session

import "github.com/gin-gonic/gin"

type SessionUseCase interface {
	SessionCommon

	CreateSession(ctx *gin.Context, username, token, timestamp string) error
}
