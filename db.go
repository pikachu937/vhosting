package vh

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mikerumy/vhosting/internal/config"
	"github.com/sirupsen/logrus"
)

const (
	connTimeOut = 5
)

func NewDBConnection(cfg config.DBConfig) *sqlx.DB {
	timeAtStarting := time.Now()

	var db *sqlx.DB

	go func() *sqlx.DB {
		for {
			db, _ = sqlx.Open(cfg.DBDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
				cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

			time.Sleep(3 * time.Millisecond)

			if db.Ping() == nil {
				return db
			}
		}
	}()

	time.Sleep(5 * time.Millisecond)

	for t := connTimeOut; t > 0; t-- {
		if db != nil {
			logrus.Printf("opened connection to db (in %s)\n", time.Since(timeAtStarting))
			return db
		}

		time.Sleep(time.Second)
	}

	logrus.Panicf("time waiting of connection to db exceeded limit (%d)\n", connTimeOut)
	return nil
}

func CloseDBConnection(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		logrus.Panicf("error occured on db connection close: %s\n", err.Error())
	}

	logrus.Printf("closed connection to db\n")
}
