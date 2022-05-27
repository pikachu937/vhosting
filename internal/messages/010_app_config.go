package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func WarningCannotConvertCvar(cvarName string, setValue interface{}, err error) *lg.Log {
	return &lg.Log{ErrCode: 10, Message: fmt.Sprintf("Cannot convert cvar %s. Set default value: %v. Error: %s.", cvarName, setValue, err.Error()), ErrLevel: logger.ErrLevelWarning}
}

func FatalFailedToLoadConfigFile(err error) *lg.Log {
	return &lg.Log{ErrCode: 11, Message: fmt.Sprintf("Failed to load config file. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelFatal}
}

func InfoConfigLoaded() *lg.Log {
	return &lg.Log{Message: "Config loaded.", ErrLevel: logger.ErrLevelInfo}
}
