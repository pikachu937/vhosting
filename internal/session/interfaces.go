package session

import "github.com/gin-gonic/gin"

type SessCommon interface {
	DeleteSession(token string) error
	IsSessionExists(token string) (bool, error)
	GetSessionAndDate(token string) (*Session, error)
}

type SessUseCase interface {
	SessCommon

	CreateSession(ctx *gin.Context, username, token, timestamp string) error
}

type SessRepository interface {
	SessCommon

	CreateSession(session Session) error
}
