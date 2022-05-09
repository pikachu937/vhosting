package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorTimeWaitingOfDBConnectionExceededLimit(timeout int) *models.Log {
	return &models.Log{ErrorCode: 30, Message: fmt.Sprintf("Time waiting of DB connection exceeded limit (%d seconds).", timeout), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCloseDBConnection(err error) *models.Log {
	return &models.Log{ErrorCode: 31, Message: fmt.Sprintf("Cannot close DB connection. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}
