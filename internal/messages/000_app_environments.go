package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func FatalFailedToLoadEnvironmentFile(err error) *lg.Log {
	return &lg.Log{Message: "Failed to load environment file. Error: " + err.Error(), ErrLevel: logger.ErrLevelFatal}
}

func InfoEnvironmentsLoaded() *lg.Log {
	return &lg.Log{Message: "Environments loaded"}
}
