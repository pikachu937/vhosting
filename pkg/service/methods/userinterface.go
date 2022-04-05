package service

import (
	vhs "github.com/mikerumy/vhservice"
	repository "github.com/mikerumy/vhservice/pkg/repository/interfaces"
)

type UserInterfaceService struct {
	repo repository.UserInterface
}

func NewUserInterfaceService(repo repository.UserInterface) *UserInterfaceService {
	return &UserInterfaceService{repo: repo}
}

func (s *UserInterfaceService) POSTUser(user vhs.User) (int, error) {
	return s.repo.POSTUser(user)
}

func (s *UserInterfaceService) GETUser(id int) (*vhs.User, error) {
	return s.repo.GETUser(id)
}

func (s *UserInterfaceService) GETAllUsers() (map[int]*vhs.User, error) {
	return s.repo.GETAllUsers()
}

func (s *UserInterfaceService) PUTUser(id int, user vhs.User) (int, error) {
	return s.repo.PUTUser(id, user)
}

func (s *UserInterfaceService) PATCHUser(id int, user vhs.User) (int, error) {
	return s.repo.PATCHUser(id, user)
}

func (s *UserInterfaceService) DELETEUser(id int) (int, error) {
	return s.repo.DELETEUser(id)
}
