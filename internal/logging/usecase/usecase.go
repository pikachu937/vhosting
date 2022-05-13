package usecase

import (
	"errors"
	"fmt"
	"reflect"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

type LogUseCase struct {
	logRepo lg.LogRepository
}

func NewLogUseCase(logRepo lg.LogRepository) *LogUseCase {
	return &LogUseCase{
		logRepo: logRepo,
	}
}

func (u *LogUseCase) CreateLogRecord(log *lg.Log) error {
	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		return u.logRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUser {
		log.Message = logger.GotUserData
		return u.logRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUsersMap {
		log.Message = logger.GotAllUsersData
		return u.logRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUserpermsMap {
		log.Message = logger.GotUserPermissions
		return u.logRepo.CreateLogRecord(log)
	}
	return errors.New("Undefined type of message.")
}
