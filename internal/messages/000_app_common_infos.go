package messages

import (
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoEnvironmentVarsLoaded() *models.Log {
	return &models.Log{Message: "Environment vars loaded.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoConfigVarsLoaded() *models.Log {
	return &models.Log{Message: "Config vars loaded.", ErrorLevel: logger.ErrLevelInfo}
}
