package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/config_tool"
)

type SessionUseCase struct {
	cfg         config_tool.Config
	sessionRepo session.SessionRepository
	authRepo    auth.AuthRepository
}

func NewSessionUseCase(cfg config_tool.Config, sessionRepo session.SessionRepository,
	authRepo auth.AuthRepository) *SessionUseCase {
	return &SessionUseCase{
		cfg:         cfg,
		sessionRepo: sessionRepo,
		authRepo:    authRepo,
	}
}

func (u *SessionUseCase) IsSessionExists(token string) (bool, error) {
	return u.sessionRepo.IsSessionExists(token)
}

func (u *SessionUseCase) CreateSession(ctx *gin.Context, username, token, timestamp string) error {
	var sess session.Session
	var err error
	sess.Content = token
	sess.CreationDate = timestamp
	if err != nil {
		return err
	}

	if err = u.sessionRepo.CreateSession(sess); err != nil {
		return err
	}

	return u.authRepo.UpdateNamepassLastLogin(username, sess.CreationDate)
}

func (u *SessionUseCase) DeleteSession(token string) error {
	return u.sessionRepo.DeleteSession(token)
}
