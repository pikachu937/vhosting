package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mikerumy/vhosting2/internal/config"
	"github.com/mikerumy/vhosting2/internal/models"
	"github.com/mikerumy/vhosting2/pkg/logger"
	"github.com/mikerumy/vhosting2/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set up logger
	if err := logger.InitLogger(); err != nil {
		logrus.Warningf("Cannot initialize logger. Error: %s.\n", err.Error())
	}

	// Load environment file
	if err := godotenv.Load("./configs/.env"); err != nil {
		logrus.Fatalf("Failed to load environment file. Error: %s.\n", err.Error())
	}

	var cfg models.Config
	var err error

	// Load config file
	if cfg, err = config.LoadConfig("./configs/config.yml"); err != nil {
		logrus.Fatalf("Failed to load config file. Error: %s.\n", err.Error())
	}

	app := server.NewApp(cfg)

	// Run the server and wait for ^C signal from keyboard to shut down
	if err := app.Run(); err != nil {
		logrus.Fatalf("Failure on server running. Error: %s.\n", err.Error())
	}
}
