package service

import vhs "github.com/mikerumy/vhservice"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserInterface interface {
	POSTUser(user vhs.User) (int, error)
	GETUser(id int) (*vhs.User, error)
	GETAllUsers() (map[int]*vhs.User, error)
	PUTUser(id int, user vhs.User) (int, error)
	PATCHUser(id int, user vhs.User) (int, error)
	DELETEUser(id int) (int, error)
}
