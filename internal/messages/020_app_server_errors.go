package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func WarningCannotGetLocalIP(err error) *models.Log {
	return &models.Log{ErrorCode: 20, Message: fmt.Sprintf("Cannot get local IP. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelWarning}
}
