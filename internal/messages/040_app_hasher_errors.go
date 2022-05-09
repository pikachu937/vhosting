package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotWriteBytesIntoInternalVariable(err error) *models.Log {
	return &models.Log{ErrorCode: 40, Message: fmt.Sprintf("Cannot write bytes into internal variable. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}
