package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadStreamConfigFile(err error) *lg.Log {
	return &lg.Log{ErrCode: 15, Message: "Failed to load stream config file. Error: " + err.Error(), ErrLevel: logger.ErrLevelFatal}
}

func InfoStreamConfigLoaded() *lg.Log {
	return &lg.Log{Message: "Stream config loaded"}
}
