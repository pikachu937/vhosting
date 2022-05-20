package messages

import (
	"fmt"
	"time"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoEstablishedOpenedDBConnection(timeSinceOpen time.Time) *lg.Log {
	return &lg.Log{Message: fmt.Sprintf("Established opening of connection to DB in %s.", time.Since(timeSinceOpen).Round(time.Millisecond).String()), ErrLevel: logger.ErrLevelInfo}
}

func ErrorTimeWaitingOfDBConnectionExceededLimit(timeout int) *lg.Log {
	return &lg.Log{ErrCode: 30, Message: fmt.Sprintf("Time waiting of DB connection exceeded limit (%d seconds).", timeout), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCloseDBConnection(err error) *lg.Log {
	return &lg.Log{ErrCode: 31, Message: fmt.Sprintf("Cannot close DB connection. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoEstablishedClosedConnectionToDB() *lg.Log {
	return &lg.Log{Message: "Established closing of connection to DB.", ErrLevel: logger.ErrLevelInfo}
}
