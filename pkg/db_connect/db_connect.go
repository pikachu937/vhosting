package db_connect

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/config"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func NewDBConnection(cfg *config.Config) *sqlx.DB {
	timeAtStarting := time.Now()

	var db *sqlx.DB

	sslMode := "disable"
	if cfg.DBSSLEnable {
		sslMode = "enable"
	}

	go func() *sqlx.DB {
		for {
			db, _ = sqlx.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
				cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBName, cfg.DBPassword, sslMode))

			time.Sleep(3 * time.Millisecond)

			if db.Ping() == nil {
				return db
			}
		}
	}()

	connLatency := time.Duration(cfg.DBConnectionLatencyMilliseconds)
	time.Sleep(connLatency * time.Millisecond)

	connTimeout := cfg.DBConnectionTimeoutSeconds
	for t := connTimeout; t > 0; t-- {
		if db != nil {
			if cfg.DBConnectionShowStatus {
				logger.Print(msg.InfoEstablishedOpenedDBConnection(timeAtStarting))
				return db
			}
			return db
		}
		time.Sleep(time.Second)
	}

	logger.Print(msg.ErrorTimeWaitingOfDBConnectionExceededLimit(connTimeout))
	return nil
}

func CloseDBConnection(cfg *config.Config, db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logger.Print(msg.ErrorCannotCloseDBConnection(err))
		return
	}

	if cfg.DBConnectionShowStatus {
		logger.Print(msg.InfoEstablishedClosedConnectionToDB())
		return
	}
}
