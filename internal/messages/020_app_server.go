package messages

import (
	"fmt"
	"os"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoServerStartedSuccessfullyAtLocalAddress(host, port string) *lg.Log {
	return &lg.Log{Message: fmt.Sprintf("Server started successfully at local address: %s:%s.", host, port), ErrLevel: logger.ErrLevelInfo}
}

func InfoServerShutedDownCorrectly() *lg.Log {
	return &lg.Log{Message: "Server shuted down correctly.", ErrLevel: logger.ErrLevelInfo}
}

func FatalFailureOnServerRunning(err error) *lg.Log {
	return &lg.Log{ErrCode: 20, Message: fmt.Sprintf("Failure on server running. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelFatal}
}

func WarningCannotGetLocalIP(err error) *lg.Log {
	return &lg.Log{ErrCode: 21, Message: fmt.Sprintf("Cannot get local IP. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelWarning}
}

func InfoRecivedSignal(signal os.Signal) *lg.Log {
	return &lg.Log{Message: fmt.Sprintf("Recived signal: %s.", signal), ErrLevel: logger.ErrLevelInfo}
}
