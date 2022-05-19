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
	messageType := fmt.Sprintf("%T", log.Message)
	if messageType == logger.TypeOfUser {
		log.Message = logger.GotUser
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfUsers {
		log.Message = logger.GotAllUsers
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfGroup {
		log.Message = logger.GotGroup
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfGroups {
		log.Message = logger.GotAllGroups
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfPermIds {
		log.Message = logger.GotUserPerms
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfPerms {
		log.Message = logger.GotAllPerms
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfInfo {
		log.Message = logger.GotInfo
		return u.logRepo.CreateLogRecord(log)
	}
	if messageType == logger.TypeOfInfos {
		log.Message = logger.GotAllInfos
		return u.logRepo.CreateLogRecord(log)
	}
	return errors.New("Undefined type of message. Type: " + messageType)
}
