package usecase

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type LogUseCase struct {
	logRepo logger.LogRepository
}

func NewLogUseCase(logRepo logger.LogRepository) *LogUseCase {
	return &LogUseCase{
		logRepo: logRepo,
	}
}

func (u *LogUseCase) Report(ctx *gin.Context, log *logger.Log, messageLog *logger.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := u.logRepo.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (u *LogUseCase) ReportWithToken(ctx *gin.Context, log *logger.Log, messageLog *logger.Log, token string) {
	logger.Complete(log, messageLog)
	responder.ResponseToken(ctx, log, token)
	if err := u.logRepo.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.ResponseToken(ctx, log, token)
	}
	logger.Print(log)
}

func (u *LogUseCase) CreateLogRecord(log *logger.Log) error {
	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		return u.logRepo.CreateLogRecord(log)
	}
	messageType := fmt.Sprintf("%T", log.Message)
	if messageType == logger.TypeOfUser {
		log.Message = logger.GotUser
	} else if messageType == logger.TypeOfUsers {
		log.Message = logger.GotAllUsers
	} else if messageType == logger.TypeOfGroup {
		log.Message = logger.GotGroup
	} else if messageType == logger.TypeOfGroups {
		log.Message = logger.GotAllGroups
	} else if messageType == logger.TypeOfPermIds {
		log.Message = logger.GotUserPerms
	} else if messageType == logger.TypeOfPerms {
		log.Message = logger.GotAllPerms
	} else if messageType == logger.TypeOfInfo {
		log.Message = logger.GotInfo
	} else if messageType == logger.TypeOfInfos {
		log.Message = logger.GotAllInfos
	} else if messageType == logger.TypeOfVideo {
		log.Message = logger.GotVideo
	} else if messageType == logger.TypeOfVideos {
		log.Message = logger.GotAllVideos
	} else if messageType == logger.TypeOfGroupIds {
		log.Message = logger.GotUserGroups
	} else if messageType == logger.TypeOfDownload {
		log.Message = logger.GotDownload
	} else {
		return errors.New("Undefined type of message. Type: " + messageType)
	}
	return u.logRepo.CreateLogRecord(log)
}
