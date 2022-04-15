package service

import (
	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/internal/session"
)

type Authorization interface {
	POSTSession(sess session.Session) error
	GETNamePass(namepass vh.NamePass) error
	DELETECurrentSession(cookieValue string) error
	UPDATELoginTimestamp(username, timestamp string) error
	UPDATEUserPassword(namepass vh.NamePass) error
}
