package messages

import (
	"fmt"
	"time"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoEstablishedOpenedDBConnection(timeSinceOpen time.Time) *logging.Log {
	return &logging.Log{Message: fmt.Sprintf("Established opening of connection to DB in %s.", time.Since(timeSinceOpen).Round(time.Millisecond).String()), ErrorLevel: logger.ErrLevelInfo}
}

func ErrorTimeWaitingOfDBConnectionExceededLimit(timeout int) *logging.Log {
	return &logging.Log{ErrorCode: 30, Message: fmt.Sprintf("Time waiting of DB connection exceeded limit (%d seconds).", timeout), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCloseDBConnection(err error) *logging.Log {
	return &logging.Log{ErrorCode: 31, Message: fmt.Sprintf("Cannot close DB connection. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoEstablishedClosedConnectionToDB() *logging.Log {
	return &logging.Log{Message: "Established closing of connection to DB.", ErrorLevel: logger.ErrLevelInfo}
}
