package auth

import "github.com/mikerumy/vhosting2/internal/models"

type AuthRepository interface {
	AuthCommon

	CreateSession(sess models.Session) error
	UpdateLoginTimestamp(username, token string) error
}
