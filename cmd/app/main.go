package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mikerumy/vhosting/internal/config"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
	"github.com/mikerumy/vhosting/pkg/server"
)

func main() {
	// Load environment file
	if err := godotenv.Load("./configs/.env"); err != nil {
		response.FatalFailedLoadEnvironment(err)
		return
	}
	response.InfoEnvironmentVarsLoaded()

	var cfg models.Config
	var err error

	// Load config file
	if cfg, err = config.LoadConfig("./configs/config.yml"); err != nil {
		response.FatalFailedLoadConfig(err)
		return
	}
	response.InfoConfigVarsLoaded()

	app := server.NewApp(cfg)

	// Run the server and wait for ^C signal from keyboard to shut down
	if err := app.Run(); err != nil {
		response.FatalFailureOnServerRunning(err)
	}
}
