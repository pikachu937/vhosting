package storage

import (
	vhs "github.com/mikerumy/vhservice"
	interfaces "github.com/mikerumy/vhservice/pkg/storage/interfaces"
	methods "github.com/mikerumy/vhservice/pkg/storage/methods"
)

type Storage struct {
	interfaces.UserInterface
	interfaces.Authorization
}

func NewStorage(cfg vhs.DBConfig) *Storage {
	return &Storage{
		UserInterface: methods.NewUserInterfaceStorage(cfg),
		Authorization: methods.NewAuthorizationStorage(cfg),
	}
}
