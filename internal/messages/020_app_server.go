package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoServerWasSuccessfullyStartedAtLocalIP(host, port string) *logging.Log {
	return &logging.Log{Message: fmt.Sprintf("Server was successfully started at local IP: %s:%s.", host, port), ErrorLevel: logger.ErrLevelInfo}
}

func InfoServerWasGracefullyShutDown() *logging.Log {
	return &logging.Log{Message: "Server was gracefully shut down.", ErrorLevel: logger.ErrLevelInfo}
}

func FatalFailureOnServerRunning(err error) *logging.Log {
	return &logging.Log{ErrorCode: 20, Message: fmt.Sprintf("Failure on server running. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelFatal}
}

func WarningCannotGetLocalIP(err error) *logging.Log {
	return &logging.Log{ErrorCode: 21, Message: fmt.Sprintf("Cannot get local IP. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelWarning}
}
