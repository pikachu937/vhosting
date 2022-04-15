package hashing

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikerumy/vhosting/internal/user"
)

const (
	salt       = "jK@s13DvU3o3H#e0N7j9G@h9K7r#Ps"
	signingKey = "jD2@hSw2eGe7#HkU7fH@8kLe0#6GeD"
	tokenTTL   = 24 * time.Hour
)

func GeneratePasswordHash(password string) string {
	if password == "" {
		return ""
	}

	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		username,
		password,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (user.NamePass, error) {
	var namepass user.NamePass
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return namepass, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return namepass, errors.New("token claims are not of type *tokenClaims")
	}

	namepass.Username = claims.Username
	namepass.PasswordHash = claims.Password
	return namepass, nil
}
