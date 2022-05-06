package hashing

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func ErrorCannotWriteBytesIntoHashingVariable(err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  501,
		Message: fmt.Sprintf("Cannot write bytes into hashing variable. Error: %s.",
			err.Error()),
	})
}
