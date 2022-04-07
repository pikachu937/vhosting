package service

import vh "github.com/mikerumy/vhosting"

type Authorization interface {
	POSTUser(user vh.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
