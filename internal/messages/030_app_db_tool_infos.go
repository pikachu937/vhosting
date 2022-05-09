package messages

import (
	"fmt"
	"time"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoEstablishedOpenedDBConnection(timeSinceOpen time.Time) *models.Log {
	return &models.Log{Message: fmt.Sprintf("Established opening of connection to DB in %s.", time.Since(timeSinceOpen).Round(time.Millisecond).String()), ErrorLevel: logger.ErrLevelInfo}
}

func InfoEstablishedClosedConnectionToDB() *models.Log {
	return &models.Log{Message: "Established closing of connection to DB.", ErrorLevel: logger.ErrLevelInfo}
}
