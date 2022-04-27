package auth

import "github.com/mikerumy/vhosting2/internal/models"

type AuthCommon interface {
	GetNamepass(namepass models.Namepass) error
	DeleteSession(token string) error
	UpdateUserPassword(namepass models.Namepass) error
	IsNamepassExists(username, passwordHash string) (bool, error)
	IsSessionExists(token string) (bool, error)
}
