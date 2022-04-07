package storage

import vh "github.com/mikerumy/vhosting"

type Authorization interface {
	POSTUser(user vh.User) (int, error)
	GETUser(username, password string) (vh.User, error)
}
