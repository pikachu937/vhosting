package service

import vhs "github.com/mikerumy/vhservice"

type Authorization interface {
	POSTUser(user vhs.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
