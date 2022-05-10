package usecase

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	"github.com/mikerumy/vhosting/pkg/logger"
)

type LoggingUseCase struct {
	cfg         config_tool.Config
	loggingRepo logging.LoggingRepository
}

func NewLoggingUseCase(cfg config_tool.Config, loggingRepo logging.LoggingRepository) *LoggingUseCase {
	return &LoggingUseCase{
		cfg:         cfg,
		loggingRepo: loggingRepo,
	}
}

func (u *LoggingUseCase) CreateLogRecord(log *logging.Log) error {
	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		return u.loggingRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUser {
		log.Message = logger.GotUserData
		return u.loggingRepo.CreateLogRecord(log)
	}
	if fmt.Sprintf("%T", log.Message) == logger.TypeUsersSlice {
		log.Message = logger.GotAllUsersData
		return u.loggingRepo.CreateLogRecord(log)
	}
	return errors.New("Undefined type of message.")
}
