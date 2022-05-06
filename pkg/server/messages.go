package server

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func InfoServerWasSuccessfullyStartedAtLocalIP(host, port string) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelInfo,
		Message: fmt.Sprintf("Server was successfully started at local IP: %s:%s.",
			host, port),
	})
}

func InfoServerWasGracefullyShutDown() {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelInfo,
		Message:    "Server was gracefully shut down.",
	})
}

func WarningCannotGetLocalIP(err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelWarning,
		ErrorCode:  201,
		Message: fmt.Sprintf("Cannot get local IP. Error: %s.",
			err.Error()),
	})
}
