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

const timeWaiting = 5

func NewPostgresConnection(cfg Config) *sqlx.DB {
	timeAtStarting := time.Now()

	var db *sqlx.DB
	var err error

	go func(cfg Config) {
		for {
			db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
				cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
			if err == nil {
				return
			}

			time.Sleep(3 * time.Millisecond)

			err = db.Ping()
			if err == nil {
				return
			}
		}
	}(cfg)

	time.Sleep(10 * time.Millisecond)

	if db == nil {
		for t := timeWaiting; t > 0; t-- {
			if db != nil {
				break
			}
			time.Sleep(time.Second)
		}
	}

	if db != nil {
		logrus.Printf("successfully connected to db in %s", time.Since(timeAtStarting))
		return db
	}

	logrus.Panicf("time waiting of connection to db exceeded limit (%d)\n", timeWaiting)
	return nil
}
