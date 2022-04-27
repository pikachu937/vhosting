package hashing

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikerumy/vhosting2/internal/models"
)

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(namepass models.Namepass, signingKey string, tokenTTLHours int) (string, error) {
	tokenTTL := time.Duration(tokenTTLHours) * time.Hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		namepass.Username,
		namepass.PasswordHash,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(tokenContent, signingKey string) (models.Namepass, error) {
	var namepass models.Namepass
	token, err := jwt.ParseWithClaims(tokenContent, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid signing method.")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return namepass, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return namepass, errors.New("Token claims are not of type \"*tokenClaims\".")
	}

	namepass.Username = claims.Username
	namepass.PasswordHash = claims.Password
	return namepass, nil
}
