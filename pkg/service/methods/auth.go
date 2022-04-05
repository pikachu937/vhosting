package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	vhs "github.com/mikerumy/vhservice"
	storage "github.com/mikerumy/vhservice/pkg/storage/interfaces"
)

const (
	salt       = "jK@s13DvU3o3H#e0N7j9G@h9K7r#Ps"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	stor storage.Authorization
}

func NewAuthService(stor storage.Authorization) *AuthService {
	return &AuthService{stor: stor}
}

func (s *AuthService) POSTUser(user vhs.User) (int, error) {
	user.PasswordHash = GeneratePasswordHash(user.PasswordHash)
	return s.stor.POSTUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.stor.GETUser(username, GeneratePasswordHash(password))
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

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
