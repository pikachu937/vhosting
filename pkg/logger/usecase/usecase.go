package usecase

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

type LogUseCase struct {
	cfg     models.Config
	logRepo logger.LogRepository
}

func NewLogUseCase(cfg models.Config, logRepo logger.LogRepository) *LogUseCase {
	return &LogUseCase{
		cfg:     cfg,
		logRepo: logRepo,
	}
}

func (u *LogUseCase) CreateLogRecord(log *models.Log) error {
	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		return u.logRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUser {
		log.Message = logger.GotUserData
		return u.logRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUsersSlice {
		log.Message = logger.GotAllUsersData
		return u.logRepo.CreateLogRecord(log)
	}
	return errors.New("Undefined type of message.")
}
