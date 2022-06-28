package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/config"
	sconfig "github.com/mikerumy/vhosting/pkg/config_stream"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/server"
)

func main() {
	// Assign environments path
	if err := godotenv.Load("./configs/.env"); err != nil {
		logger.Print(msg.FatalFailedToLoadEnvironmentFile(err))
		return
	}
	logger.Print(msg.InfoEnvironmentsLoaded())

	// Load config
	cfg, err := config.LoadConfig("./configs/config.yml")
	if err != nil {
		logger.Print(msg.FatalFailedToLoadConfigFile(err))
		return
	}
	logger.Print(msg.InfoConfigLoaded())

	// Create stream config
	var scfg sconfig.Config

	// Init new server
	app := server.NewApp(cfg, &scfg)

	// Start tasks
	runTasks(app)

	// Start server, wait interrupt
	if err := app.Run(); err != nil {
		logger.Print(msg.FatalFailureOnServerRunning(err))
	}
}
