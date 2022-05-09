package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	logger "github.com/mikerumy/vhosting/pkg/logger"
)

func InfoServerWasSuccessfullyStartedAtLocalIP(host, port string) *models.Log {
	return &models.Log{Message: fmt.Sprintf("Server was successfully started at local IP: %s:%s.", host, port), ErrorLevel: logger.ErrLevelInfo}
}

func InfoServerWasGracefullyShutDown() *models.Log {
	return &models.Log{Message: "Server was gracefully shut down.", ErrorLevel: logger.ErrLevelInfo}
}
