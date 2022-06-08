package messages

import (
	"fmt"
	"os"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoServerStartedSuccessfullyAtLocalAddress(host string, port int) *lg.Log {
	return &lg.Log{Message: "Server was successfully started at local address: " + fmt.Sprintf("%s:%d", host, port)}
}

func InfoServerShutedDownCorrectly() *lg.Log {
	return &lg.Log{Message: "Server was correctly shuted down"}
}

func FatalFailureOnServerRunning(err error) *lg.Log {
	return &lg.Log{ErrCode: 20, Message: "Failure on server running. Error: " + err.Error(), ErrLevel: logger.ErrLevelFatal}
}

func WarningCannotGetLocalIP(err error) *lg.Log {
	return &lg.Log{ErrCode: 21, Message: "Cannot get local IP. Error: " + err.Error(), ErrLevel: logger.ErrLevelWarning}
}

func InfoRecivedSignal(signal os.Signal) *lg.Log {
	return &lg.Log{Message: "Recived signal: " + signal.String()}
}
