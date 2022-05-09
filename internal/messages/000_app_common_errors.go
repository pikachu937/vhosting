package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadEnvironmentFile(err error) *models.Log {
	return &models.Log{ErrorCode: 0, Message: fmt.Sprintf("Failed to load environment file. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func FatalFailedToLoadConfigFile(err error) *models.Log {
	return &models.Log{ErrorCode: 1, Message: fmt.Sprintf("Failed to load config file. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func FatalFailureOnServerRunning(err error) *models.Log {
	return &models.Log{ErrorCode: 2, Message: fmt.Sprintf("Failure on server running. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}
