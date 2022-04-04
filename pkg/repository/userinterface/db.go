package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	dbDriver    = "postgres"
	connTimeOut = 5
)

func NewDBConnection(cfg Config) *sqlx.DB {
	timeAtStarting := time.Now()

	var db *sqlx.DB

	go func() *sqlx.DB {
		for {
			db, _ = sqlx.Open(dbDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
				cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

			time.Sleep(3 * time.Millisecond)

			if db.Ping() == nil {
				return db
			}
		}
	}()

	time.Sleep(10 * time.Millisecond)

	for t := connTimeOut; t > 0; t-- {
		if db != nil {
			logrus.Printf("successfully connected to db in %s", time.Since(timeAtStarting))
			return db
		}

		time.Sleep(time.Second)
	}

	logrus.Panicf("time waiting of connection to db exceeded limit (%d)\n", connTimeOut)
	return nil
}
