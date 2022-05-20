package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoServerWasSuccessfullyStartedAtLocalIP(host, port string) *lg.Log {
	return &lg.Log{Message: fmt.Sprintf("Server was successfully started at local IP: %s:%s.", host, port), ErrLevel: logger.ErrLevelInfo}
}

func InfoServerWasGracefullyShutDown() *lg.Log {
	return &lg.Log{Message: "Server was correctly shuted down.", ErrLevel: logger.ErrLevelInfo}
}

func FatalFailureOnServerRunning(err error) *lg.Log {
	return &lg.Log{ErrCode: 20, Message: fmt.Sprintf("Failure on server running. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelFatal}
}

func WarningCannotGetLocalIP(err error) *lg.Log {
	return &lg.Log{ErrCode: 21, Message: fmt.Sprintf("Cannot get local IP. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelWarning}
}
