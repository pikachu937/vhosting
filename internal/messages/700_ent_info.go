package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/info"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorStreamCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 700, Message: fmt.Sprintf("Stream cannot be empty."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateInfo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 701, Message: fmt.Sprintf("Cannot create info. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoInfoCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Info created.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotCheckInfoExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 702, Message: fmt.Sprintf("Cannot check info existence. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorInfoWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 703, Message: fmt.Sprintf("Info with requested ID is not exist."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGetInfo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 704, Message: fmt.Sprintf("Cannot get info. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGotInfo(nfo *info.Info) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: nfo, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllInfos(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 705, Message: fmt.Sprintf("Cannot get all infos. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoNoInfosAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No infos available.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotAllInfos(users map[int]*info.Info) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotPartiallyUpdateInfo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 706, Message: fmt.Sprintf("Cannot partially update info. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoInfoPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Info partially updated.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteInfo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 707, Message: fmt.Sprintf("Cannot delete info. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoInfoDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Info deleted.", ErrorLevel: logger.ErrLevelInfo}
}
