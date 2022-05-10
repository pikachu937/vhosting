package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadEnvironmentFile(err error) *logging.Log {
	return &logging.Log{ErrorCode: 0, Message: fmt.Sprintf("Failed to load environment file. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func InfoEnvironmentVarsLoaded() *logging.Log {
	return &logging.Log{Message: "Environment vars loaded.", ErrorLevel: logger.ErrLevelInfo}
}
