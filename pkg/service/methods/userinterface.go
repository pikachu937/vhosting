package service

import (
	vh "github.com/mikerumy/vhosting"
	storage "github.com/mikerumy/vhosting/pkg/storage/interfaces"
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

func (s *UserInterfaceService) POSTUser(user vh.User) error {
	return s.stor.POSTUser(user)
}

func (s *UserInterfaceService) GETUser(id int) (*vh.User, error) {
	return s.stor.GETUser(id)
}

func (s *UserInterfaceService) GETAllUsers() (map[int]*vh.User, error) {
	return s.stor.GETAllUsers()
}

func (s *UserInterfaceService) PATCHUser(id int, user vh.User) error {
	return s.stor.PATCHUser(id, user)
}

func (s *UserInterfaceService) DELETEUser(id int) error {
	return s.stor.DELETEUser(id)
}
