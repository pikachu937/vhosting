package storage

import (
	"github.com/mikerumy/vhosting/internal/config"
	interfaces "github.com/mikerumy/vhosting/pkg/storage/interfaces"
	methods "github.com/mikerumy/vhosting/pkg/storage/methods"
)

type Storage struct {
	interfaces.UserInterface
	interfaces.Authorization
}

func NewStorage(cfg config.DBConfig) *Storage {
	return &Storage{
		UserInterface: methods.NewUserInterfaceStorage(cfg),
		Authorization: methods.NewAuthorizationStorage(cfg),
	}
}
