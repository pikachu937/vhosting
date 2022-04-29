package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func NewDBConnection(cfg models.Config) *sqlx.DB {
	timeAtStarting := time.Now()

	var db *sqlx.DB

	go func() *sqlx.DB {
		for {
			db, _ = sqlx.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
				cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode))

			time.Sleep(3 * time.Millisecond)

			if db.Ping() == nil {
				return db
			}
		}
	}()

	connLatency := time.Duration(cfg.DBConnLatencyMilliseconds)
	time.Sleep(connLatency * time.Millisecond)

	connTimeout := cfg.DBConnTimeoutSeconds
	for t := connTimeout; t > 0; t-- {
		if db != nil {
			if cfg.DBLogConnStatus {
				response.Response(nil, models.Log{Message: fmt.Sprintf("Established opened DB connection in %s.",
					time.Since(timeAtStarting).Round(time.Millisecond).String())})
				return db
			}

			return db
		}

		time.Sleep(time.Second)
	}

	response.Response(nil, models.Log{Message: fmt.Sprintf("Time waiting of DB connection exceeded limit (%d seconds).", connTimeout)})
	return nil
}

func CloseDBConnection(cfg models.Config, db *sqlx.DB) {
	if err := db.Close(); err != nil {
		response.Response(nil, models.Log{Message: fmt.Sprintf("Cannot close DB connection. Error: %s.", err.Error())})
	}

	if cfg.DBLogConnStatus {
		response.Response(nil, models.Log{Message: "Established closed connection to DB."})
		return
	}
}
