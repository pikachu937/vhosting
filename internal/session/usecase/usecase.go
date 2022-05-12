package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	sess "github.com/mikerumy/vhosting/internal/session"
)

type SessUseCase struct {
	sessRepo sess.SessRepository
	authRepo auth.AuthRepository
}

func NewSessUseCase(sessRepo sess.SessRepository,
	authRepo auth.AuthRepository) *SessUseCase {
	return &SessUseCase{
		sessRepo: sessRepo,
		authRepo: authRepo,
	}
}

func (u *SessUseCase) IsSessionExists(token string) (bool, error) {
	return u.sessRepo.IsSessionExists(token)
}

func (u *SessUseCase) CreateSession(ctx *gin.Context, username, token, timestamp string) error {
	var session sess.Session
	var err error
	session.Content = token
	session.CreationDate = timestamp
	if err != nil {
		return err
	}

	if err = u.sessRepo.CreateSession(session); err != nil {
		return err
	}

	return u.authRepo.UpdateNamepassLastLogin(username, session.CreationDate)
}

func (u *SessUseCase) DeleteSession(token string) error {
	return u.sessRepo.DeleteSession(token)
}
