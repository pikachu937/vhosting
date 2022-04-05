package service

import (
	vhs "github.com/mikerumy/vhservice"
	storage "github.com/mikerumy/vhservice/pkg/storage/interfaces"
)

type UserInterfaceService struct {
	stor storage.UserInterface
}

func NewUserInterfaceService(stor storage.UserInterface) *UserInterfaceService {
	return &UserInterfaceService{stor: stor}
}

func (s *UserInterfaceService) POSTUser(user vhs.User) (int, error) {
	return s.stor.POSTUser(user)
}

func (s *UserInterfaceService) GETUser(id int) (*vhs.User, error) {
	return s.stor.GETUser(id)
}

func (s *UserInterfaceService) GETAllUsers() (map[int]*vhs.User, error) {
	return s.stor.GETAllUsers()
}

func (s *UserInterfaceService) PUTUser(id int, user vhs.User) (int, error) {
	return s.stor.PUTUser(id, user)
}

func (s *UserInterfaceService) PATCHUser(id int, user vhs.User) (int, error) {
	return s.stor.PATCHUser(id, user)
}

func (s *UserInterfaceService) DELETEUser(id int) (int, error) {
	return s.stor.DELETEUser(id)
}
