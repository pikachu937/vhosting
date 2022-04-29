package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/cookie"
	"github.com/mikerumy/vhosting/pkg/hashing"
	"github.com/mikerumy/vhosting/pkg/timestamp"
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

func (u *AuthUseCase) CreateSession(username, token string) error {
	thisTimestamp := timestamp.WriteThisTimestamp()
	var sess models.Session
	sess.Content = token
	sess.CreationDate = thisTimestamp
	if err := u.authRepo.CreateSession(sess); err != nil {
		return err
	}

	return u.authRepo.UpdateLoginTimestamp(username, thisTimestamp)
}

func (u *AuthUseCase) ReadCookie(ctx *gin.Context) string {
	return cookie.ReadCookie(ctx)
}

func (u *AuthUseCase) BindJSONNamepass(ctx *gin.Context) (models.Namepass, error) {
	var namepass models.Namepass
	if err := ctx.BindJSON(&namepass); err != nil {
		return namepass, err
	}
	if namepass.PasswordHash != "" {
		namepass.PasswordHash = hashing.GeneratePasswordHash(namepass.PasswordHash, u.cfg.HashingPasswordSalt)
	}
	return namepass, nil
}

func (u *AuthUseCase) GenerateToken(namepass models.Namepass) (string, error) {
	namepass.PasswordHash = hashing.GeneratePasswordHash(namepass.PasswordHash, u.cfg.HashingPasswordSalt)
	token, err := hashing.GenerateToken(namepass, u.cfg.HashingTokenSigningKey, u.cfg.SessionTTLHours)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *AuthUseCase) SendCookie(ctx *gin.Context, token string) {
	cookie.SendCookie(ctx, token, u.cfg.SessionTTLHours)
}

func (u *AuthUseCase) ParseToken(token string) (models.Namepass, error) {
	return hashing.ParseToken(token, u.cfg.HashingTokenSigningKey)
}

func (u *AuthUseCase) DeleteCookie(ctx *gin.Context) {
	cookie.DeleteCookie(ctx)
}
