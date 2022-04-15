package service

import (
	storage "github.com/mikerumy/vhosting/pkg/storage/interfaces"

	user "github.com/mikerumy/vhosting/internal/user"
)

type UserInterfaceService struct {
	stor storage.UserInterface
}

func NewUserInterfaceService(stor storage.UserInterface) *UserInterfaceService {
	return &UserInterfaceService{stor: stor}
}

func (s *UserInterfaceService) CheckUserExistence(idOrUsername interface{}) (bool, error) {
	return s.stor.CheckUserExistence(idOrUsername)
}

func (s *UserInterfaceService) POSTUser(usr user.User) error {
	return s.stor.POSTUser(usr)
}

func (s *UserInterfaceService) GETUser(id int) (*user.User, error) {
	return s.stor.GETUser(id)
}

func (s *UserInterfaceService) GETAllUsers() (map[int]*user.User, error) {
	return s.stor.GETAllUsers()
}

func (s *UserInterfaceService) PATCHUser(id int, usr user.User) error {
	return s.stor.PATCHUser(id, usr)
}

func (s *UserInterfaceService) DELETEUser(id int) error {
	return s.stor.DELETEUser(id)
}
