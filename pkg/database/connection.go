package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mikerumy/vhosting2/internal/models"
	"github.com/sirupsen/logrus"
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
				logrus.Infoln("Established opening of connection to DB. Time of connection:", time.Since(timeAtStarting).Round(time.Millisecond))
				return db
			}

			return db
		}

		time.Sleep(time.Second)
	}

	logrus.Errorf("Time waiting of DB connection exceeded limit (%d seconds).\n", connTimeout)
	return nil
}

func CloseDBConnection(cfg models.Config, db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logrus.Errorln("Cannot close connection to DB. Error:", err.Error())
	}

	if cfg.DBLogConnStatus {
		logrus.Infoln("Established closing of connection to DB.")
		return
	}
}
