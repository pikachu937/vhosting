package storage

import vh "github.com/mikerumy/vhosting"

type UserInterface interface {
	POSTUser(user vh.User) (int, error)
	GETUser(id int) (*vh.User, error)
	GETAllUsers() (map[int]*vh.User, error)
	PUTUser(id int, user vh.User) (int, error)
	PATCHUser(id int, user vh.User) (int, error)
	DELETEUser(id int) (int, error)
}
