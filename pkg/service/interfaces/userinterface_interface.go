package service

import vh "github.com/mikerumy/vhosting"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserInterface interface {
	CheckUserExistence(idOrUsername interface{}) (bool, error)
	POSTUser(user vh.User) error
	GETUser(id int) (*vh.User, error)
	GETAllUsers() (map[int]*vh.User, error)
	PATCHUser(id int, user vh.User) error
	DELETEUser(id int) error
}
