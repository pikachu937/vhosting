package messages

import (
	"strconv"
	"time"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoEstablishedOpenedDBConnection(timeSinceOpen time.Time) *lg.Log {
	return &lg.Log{Message: "Established opening of connection to DB in %s" + time.Since(timeSinceOpen).Round(time.Millisecond).String()}
}

func ErrorTimeWaitingOfDBConnectionExceededLimit(timeout int) *lg.Log {
	return &lg.Log{ErrCode: 30, Message: "Time waiting of DB connection exceeded limit (" + strconv.Itoa(timeout) + " seconds)", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCloseDBConnection(err error) *lg.Log {
	return &lg.Log{ErrCode: 31, Message: "Cannot close DB connection. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoEstablishedClosedConnectionToDB() *lg.Log {
	return &lg.Log{Message: "Established closing of connection to DB"}
}
