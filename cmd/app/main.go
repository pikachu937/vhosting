package main

import (
	msg "github.com/dmitrij/vhosting/internal/messages"
	"github.com/dmitrij/vhosting/pkg/config"
	sconfig "github.com/dmitrij/vhosting/pkg/config_stream"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Assign environments path
	if err := godotenv.Load("/var/lib/configs/.env"); err != nil {
		logger.Print(msg.FatalFailedToLoadEnvironmentFile(err))
		return
	}
	logger.Print(msg.InfoEnvironmentsLoaded())

	// Load config
	cfg, err := config.LoadConfig("/var/lib/configs/config.yml")
	if err != nil {
		logger.Print(msg.FatalFailedToLoadConfigFile(err))
		return
	}
	logger.Print(msg.InfoConfigLoaded())

	// Create stream config
	var scfg sconfig.Config

	// Init new server
	app := server.NewApp(cfg, &scfg)

	// Run stream recieving
	app.StreamUC.ServeStreams()

	// Start server, wait interrupt
	if err := app.Run(); err != nil {
		logger.Print(msg.FatalFailureOnServerRunning(err))
	}
}
