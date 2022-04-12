package service

import vh "github.com/mikerumy/vhosting"

type Authorization interface {
	POSTSession(session vh.Session) error
	GETNamePass(namepass vh.NamePass) error
	DELETECurrentSession(cookieValue string) error
	UPDATELoginTimestamp(username, timestamp string) error
	UPDATEUserPassword(namepass vh.NamePass) error
}
