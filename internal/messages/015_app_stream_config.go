package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadStreamConfigFile(err error) *lg.Log {
	return &lg.Log{ErrCode: 15, Message: fmt.Sprintf("Failed to load stream config file. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelFatal}
}

func InfoStreamConfigLoaded() *lg.Log {
	return &lg.Log{Message: "Stream config loaded.", ErrLevel: logger.ErrLevelInfo}
}
