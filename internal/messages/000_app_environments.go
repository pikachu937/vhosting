package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadEnvironmentFile(err error) *lg.Log {
	return &lg.Log{ErrCode: 0, Message: fmt.Sprintf("Failed to load environment file. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelFatal}
}

func InfoEnvironmentsLoaded() *lg.Log {
	return &lg.Log{Message: "Environments loaded.", ErrLevel: logger.ErrLevelInfo}
}
