package service

import (
	"github.com/mikerumy/vhosting/internal/session"
	user "github.com/mikerumy/vhosting/internal/user"
)

type Authorization interface {
	POSTSession(sess session.Session) error
	GETNamePass(namepass user.NamePass) error
	DELETECurrentSession(cookieValue string) error
	UPDATELoginTimestamp(username, timestamp string) error
	UPDATEUserPassword(namepass user.NamePass) error
}
