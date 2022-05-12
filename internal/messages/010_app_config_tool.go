package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func WarningCannotConvertCvar(cvarName string, setValue interface{}, err error) *lg.Log {
	return &lg.Log{ErrorCode: 10, Message: fmt.Sprintf("Cannot convert cvar %s. Set default value: %v. Error: %s.", cvarName, setValue, err.Error()), ErrorLevel: logger.ErrLevelWarning}
}

func FatalFailedToLoadConfigFile(err error) *lg.Log {
	return &lg.Log{ErrorCode: 11, Message: fmt.Sprintf("Failed to load config file. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func InfoConfigVarsLoaded() *lg.Log {
	return &lg.Log{Message: "Config vars loaded.", ErrorLevel: logger.ErrLevelInfo}
}
