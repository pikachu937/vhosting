package repository

import vhs "github.com/mikerumy/vhservice"

type Authorization interface {
	POSTUser(user vhs.User) (int, error)
	GETUser(username, password string) (vhs.User, error)
}
