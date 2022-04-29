package response

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
)

func FatalFailedLoadEnvironment(err error) {
	Response(nil, models.Log{Message: fmt.Sprintf("Failed to load environment file. Error: %s.", err.Error()), ErrorCode: 1, ErrorLevel: ErrLevelFatal})
}

func InfoEnvironmentVarsLoaded() {
	Response(nil, models.Log{Message: "Environment vars loaded."})
}

func WarningCannotConvertCvar(cvarName string, setValue interface{}, err error) {
	Response(nil, models.Log{Message: fmt.Sprintf(WarnCannotConvertCvar, cvarName, setValue, err.Error()), ErrorLevel: ErrLevelWarning})
}

func FatalFailedLoadConfig(err error) {
	Response(nil, models.Log{Message: fmt.Sprintf("Failed to load config file. Error: %s.", err.Error()), ErrorCode: 2, ErrorLevel: ErrLevelFatal})
}

func InfoConfigVarsLoaded() {
	Response(nil, models.Log{Message: "Config vars loaded."})
}

func InfoServerWasSuccessfullyStarted(host, port string) {
	Response(nil, models.Log{Message: fmt.Sprintf("Server was successfully started at local IP: %s:%s.", host, port)})
}

func InfoServerWasGracefullyShutDown() {
	Response(nil, models.Log{Message: "Server was gracefully shut down."})
}

func FatalFailureOnServerRunning(err error) {
	Response(nil, models.Log{Message: fmt.Sprintf("Failure on server running. Error: %s.", err.Error()), ErrorCode: 3, ErrorLevel: ErrLevelFatal})
}

func WarningCannotGetLocalIP(err error) {
	Response(nil, models.Log{Message: fmt.Sprintf("Cannot get local IP. Error: %s.", err.Error()), ErrorLevel: ErrLevelWarning})
}

func ErrorWriteBytesHashingVariable(err error) {
	Response(nil, models.Log{Message: fmt.Sprintf("Cannot write bytes into hashing variable. Error: %s.", err.Error()), ErrorCode: 4, ErrorLevel: ErrLevelError})
}
