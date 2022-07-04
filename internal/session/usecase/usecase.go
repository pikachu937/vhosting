package usecase

import (
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/gin-gonic/gin"
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

func (u *SessUseCase) CreateSession(ctx *gin.Context, username, token, timestamp string) error {
	var session sess.Session
	session.Content = token
	session.CreationDate = timestamp
	if err := u.sessRepo.CreateSession(&session); err != nil {
		return err
	}
	return u.authRepo.UpdateNamepassLastLogin(username, session.CreationDate)
}

func (u *SessUseCase) GetSessionAndDate(token string) (*sess.Session, error) {
	return u.sessRepo.GetSessionAndDate(token)
}

func (u *SessUseCase) DeleteSession(token string) error {
	return u.sessRepo.DeleteSession(token)
}
