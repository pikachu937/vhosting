package config

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func WarningCannotConvertCvar(cvarName string, setValue interface{}, err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelWarning,
		ErrorCode:  101,
		Message: fmt.Sprintf("Cannot convert cvar %s. Set default value: %v. Error: %s.",
			cvarName, setValue, err.Error()),
	})
}
