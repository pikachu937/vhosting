package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mikerumy/vhosting/internal/config"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/server"
)

func main() {
	// Load environment file
	if err := godotenv.Load("./configs/.env"); err != nil {
		FatalFailedToLoadEnvironmentFile(err)
		return
	}
	InfoEnvironmentVarsLoaded()

	var cfg models.Config
	var err error

	// Load config file
	if cfg, err = config.LoadConfig("./configs/config.yml"); err != nil {
		FatalFailedToLoadConfigFile(err)
		return
	}
	InfoConfigVarsLoaded()

	app := server.NewApp(cfg)

	// Run the server and wait for Ctrl+C combination from keyboard to shut down
	if err := app.Run(); err != nil {
		FatalFailureOnServerRunning(err)
	}
}
