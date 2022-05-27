package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/config"
	"github.com/mikerumy/vhosting/pkg/cookie"
	"github.com/mikerumy/vhosting/pkg/hasher"
)

type AuthUseCase struct {
	cfg      *config.Config
	authRepo auth.AuthRepository
}

func NewAuthUseCase(cfg *config.Config, authRepo auth.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		cfg:      cfg,
		authRepo: authRepo,
	}
}

func (u *AuthUseCase) IsTokenExists(token string) bool {
	if token != "" {
		return true
	}
	return false
}

func (u *AuthUseCase) IsMatched(username_1, username_2 string) bool {
	if username_1 == username_2 {
		return true
	}
	return false
}

func (u *AuthUseCase) GetNamepass(namepass auth.Namepass) error {
	return u.authRepo.GetNamepass(namepass)
}

func (u *AuthUseCase) UpdateNamepassPassword(namepass auth.Namepass) error {
	return u.authRepo.UpdateNamepassPassword(namepass)
}

func (u *AuthUseCase) IsNamepassExists(username, passwordHash string) (bool, error) {
	exists, err := u.authRepo.IsNamepassExists(username, passwordHash)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *AuthUseCase) ReadCookie(ctx *gin.Context) string {
	return cookie_manager.Read(ctx)
}

func (u *AuthUseCase) BindJSONNamepass(ctx *gin.Context) (auth.Namepass, error) {
	var namepass auth.Namepass
	if err := ctx.BindJSON(&namepass); err != nil {
		return namepass, err
	}
	if namepass.PasswordHash != "" {
		namepass.PasswordHash = hasher.GeneratePasswordHash(namepass.PasswordHash, u.cfg.HashingPasswordSalt)
	}
	return namepass, nil
}

func (u *AuthUseCase) GenerateToken(namepass auth.Namepass) (string, error) {
	namepass.PasswordHash = hasher.GeneratePasswordHash(namepass.PasswordHash, u.cfg.HashingPasswordSalt)
	token, err := hasher.GenerateToken(namepass, u.cfg.HashingTokenSigningKey, u.cfg.SessionTTLHours)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *AuthUseCase) SendCookie(ctx *gin.Context, token string) {
	cookie_manager.Send(ctx, token, u.cfg.SessionTTLHours)
}

func (u *AuthUseCase) ParseToken(token string) (auth.Namepass, error) {
	return hasher.ParseToken(token, u.cfg.HashingTokenSigningKey)
}

func (u *AuthUseCase) DeleteCookie(ctx *gin.Context) {
	cookie_manager.Delete(ctx)
}
