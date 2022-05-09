package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func WarningCannotConvertCvar(cvarName string, setValue interface{}, err error) *models.Log {
	return &models.Log{ErrorCode: 10, Message: fmt.Sprintf("Cannot convert cvar %s. Set default value: %v. Error: %s.", cvarName, setValue, err.Error()), ErrorLevel: logger.ErrLevelWarning}
}
