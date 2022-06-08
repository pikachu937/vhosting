package messages

import (
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadStreamConfigFile(err error) *logger.Log {
	return &logger.Log{ErrCode: 15, Message: "Failed to load stream config file. Error: " + err.Error(), ErrLevel: logger.ErrLevelFatal}
}

func InfoStreamConfigLoaded() *logger.Log {
	return &logger.Log{Message: "Stream config loaded"}
}
