package service

import (
	vhs "github.com/mikerumy/vhservice"
	"github.com/mikerumy/vhservice/pkg/repository"
)

type UserEditService struct {
	repo repository.UserEdit
}

func NewUserEditService(repo repository.UserEdit) *UserEditService {
	return &UserEditService{repo: repo}
}

func (s *UserEditService) CreateUser(user vhs.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserEditService) GetUser(id int) (*vhs.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserEditService) UpdateUser(id int, user vhs.User) (int, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *UserEditService) PartiallyUpdateUser(id int, user vhs.User) (int, error) {
	return s.repo.PartiallyUpdateUser(id, user)
}

func (s *UserEditService) DeleteUser(id int) (int, error) {
	return s.repo.DeleteUser(id)
}
