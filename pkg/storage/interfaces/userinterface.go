package storage

import vh "github.com/mikerumy/vhosting"

type UserInterface interface {
	POSTUser(user vh.User) error
	GETUser(id int) (*vh.User, error)
	GETAllUsers() (map[int]*vh.User, error)
	PATCHUser(id int, user vh.User) error
	DELETEUser(id int) error
}
