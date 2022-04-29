package auth

import "github.com/mikerumy/vhosting/internal/models"

type AuthRepository interface {
	AuthCommon

	CreateSession(sess models.Session) error
	UpdateLoginTimestamp(username, token string) error
}
