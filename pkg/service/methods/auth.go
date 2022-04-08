package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	vh "github.com/mikerumy/vhosting"
	storage "github.com/mikerumy/vhosting/pkg/storage/interfaces"
)

const (
	signingKey = "jD2@hSw2eGe7#HkU7fH@8kLe0#6GeD"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user-id"`
}

type AuthService struct {
	stor storage.Authorization
}

func NewAuthService(stor storage.Authorization) *AuthService {
	return &AuthService{stor: stor}
}

func (s *AuthService) POSTUser(user vh.User) (int, error) {
	user.PasswordHash = vh.GeneratePasswordHash(user.PasswordHash)
	return s.stor.POSTUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.stor.GETUser(username, vh.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
