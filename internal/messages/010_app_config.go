package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func WarningCannotConvertCvar(cvarName string, setValue interface{}) *lg.Log {
	return &lg.Log{ErrCode: 10, Message: "Cannot convert cvar " + cvarName + ". Set default value: " + setValue.(string), ErrLevel: logger.ErrLevelWarning}
}

func FatalFailedToLoadConfigFile(err error) *lg.Log {
	return &lg.Log{ErrCode: 11, Message: "Failed to load config file. Error: " + err.Error(), ErrLevel: logger.ErrLevelFatal}
}

func InfoConfigLoaded() *lg.Log {
	return &lg.Log{Message: "Config loaded"}
}
