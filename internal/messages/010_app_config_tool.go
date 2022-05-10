package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func WarningCannotConvertCvar(cvarName string, setValue interface{}, err error) *logging.Log {
	return &logging.Log{ErrorCode: 10, Message: fmt.Sprintf("Cannot convert cvar %s. Set default value: %v. Error: %s.", cvarName, setValue, err.Error()), ErrorLevel: logger.ErrLevelWarning}
}

func FatalFailedToLoadConfigFile(err error) *logging.Log {
	return &logging.Log{ErrorCode: 11, Message: fmt.Sprintf("Failed to load config file. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func InfoConfigVarsLoaded() *logging.Log {
	return &logging.Log{Message: "Config vars loaded.", ErrorLevel: logger.ErrLevelInfo}
}
