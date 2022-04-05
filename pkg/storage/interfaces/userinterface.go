package storage

import vhs "github.com/mikerumy/vhservice"

type UserInterface interface {
	POSTUser(user vhs.User) (int, error)
	GETUser(id int) (*vhs.User, error)
	GETAllUsers() (map[int]*vhs.User, error)
	PUTUser(id int, user vhs.User) (int, error)
	PATCHUser(id int, user vhs.User) (int, error)
	DELETEUser(id int) (int, error)
}
