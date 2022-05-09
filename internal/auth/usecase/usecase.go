package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/cookie_tool"
	"github.com/mikerumy/vhosting/pkg/hasher"
)

type AuthUseCase struct {
	cfg      models.Config
	authRepo auth.AuthRepository
}

func NewAuthUseCase(cfg models.Config, authRepo auth.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		cfg:      cfg,
		authRepo: authRepo,
	}
}

func (u *AuthUseCase) GetNamepass(namepass models.Namepass) error {
	return u.authRepo.GetNamepass(namepass)
}

func (u *AuthUseCase) DeleteSession(token string) error {
	return u.authRepo.DeleteSession(token)
}

func (u *AuthUseCase) UpdateUserPassword(namepass models.Namepass) error {
	return u.authRepo.UpdateUserPassword(namepass)
}

func (u *AuthUseCase) IsNamepassExists(username, passwordHash string) (bool, error) {
	exists, err := u.authRepo.IsNamepassExists(username, passwordHash)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *AuthUseCase) IsSessionExists(token string) (bool, error) {
	return u.authRepo.IsSessionExists(token)
}

func (u *AuthUseCase) CreateSession(ctx *gin.Context, username, token, timestamp string) error {
	var sess models.Session
	var err error
	sess.Content = token
	sess.CreationDate = timestamp
	if err != nil {
		return err
	}

	if err = u.authRepo.CreateSession(sess); err != nil {
		return err
	}

	return u.authRepo.UpdateLoginTimestamp(username, sess.CreationDate)
}

func (u *AuthUseCase) ReadCookie(ctx *gin.Context) string {
	return cookie_tool.Read(ctx)
}

func (u *AuthUseCase) BindJSONNamepass(ctx *gin.Context) (models.Namepass, error) {
	var namepass models.Namepass
	if err := ctx.BindJSON(&namepass); err != nil {
		return namepass, err
	}
	if namepass.PasswordHash != "" {
		namepass.PasswordHash = hasher.GeneratePasswordHash(namepass.PasswordHash, u.cfg.HashingPasswordSalt)
	}
	return namepass, nil
}

func (u *AuthUseCase) GenerateToken(namepass models.Namepass) (string, error) {
	namepass.PasswordHash = hasher.GeneratePasswordHash(namepass.PasswordHash, u.cfg.HashingPasswordSalt)
	token, err := hasher.GenerateToken(namepass, u.cfg.HashingTokenSigningKey, u.cfg.SessionTTLHours)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *AuthUseCase) SendCookie(ctx *gin.Context, token string) {
	cookie_tool.Send(ctx, token, u.cfg.SessionTTLHours)
}

func (u *AuthUseCase) ParseToken(token string) (models.Namepass, error) {
	return hasher.ParseToken(token, u.cfg.HashingTokenSigningKey)
}

func (u *AuthUseCase) DeleteCookie(ctx *gin.Context) {
	cookie_tool.Delete(ctx)
}
