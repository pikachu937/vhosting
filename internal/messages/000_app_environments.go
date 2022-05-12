package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadEnvironmentFile(err error) *lg.Log {
	return &lg.Log{ErrorCode: 0, Message: fmt.Sprintf("Failed to load environment file. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func InfoEnvironmentVarsLoaded() *lg.Log {
	return &lg.Log{Message: "Environment vars loaded.", ErrorLevel: logger.ErrLevelInfo}
}
