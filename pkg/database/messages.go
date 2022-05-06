package database

import (
	"fmt"
	"time"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func InfoEstablishedOpenedDBConnection(timeSinceOpen time.Time) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelInfo,
		Message: fmt.Sprintf("Established opening of connection to DB in %s.",
			time.Since(timeSinceOpen).Round(time.Millisecond).String()),
	})
}

func InfoEstablishedClosedConnectionToDB() {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelInfo,
		Message:    "Established closing of connection to DB.",
	})
}

func ErrorTimeWaitingOfDBConnectionExceededLimit(timeout int) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  301,
		Message: fmt.Sprintf("Time waiting of DB connection exceeded limit (%d seconds).",
			timeout),
	})
}

func ErrorCannotCloseDBConnection(err error) {
	response.Response(nil, models.Log{
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  302,
		Message: fmt.Sprintf("Cannot close DB connection. Error: %s.",
			err.Error()),
	})
}
