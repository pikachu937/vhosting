package main

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func FatalFailedToLoadEnvironmentFile(err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelFatal,
		ErrorCode:  1,
		Message: fmt.Sprintf("Failed to load environment file. Error: %s.",
			err.Error()),
	})
}

func FatalFailedToLoadConfigFile(err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelFatal,
		ErrorCode:  2,
		Message: fmt.Sprintf("Failed to load config file. Error: %s.",
			err.Error()),
	})
}

func FatalFailureOnServerRunning(err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelFatal,
		ErrorCode:  3,
		Message: fmt.Sprintf("Failure on server running. Error: %s.",
			err.Error()),
	})
}

func InfoEnvironmentVarsLoaded() {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelInfo,
		Message:    "Environment vars loaded.",
	})
}

func InfoConfigVarsLoaded() {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelInfo,
		Message:    "Config vars loaded.",
	})
}
