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
	// Load environments
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

	// Load stream config
	scfg, err := sconfig.LoadConfig("./configs/stream_config.json")
	if err != nil {
		logger.Print(msg.FatalFailedToLoadStreamConfigFile(err))
		return
	}
	logger.Print(msg.InfoStreamConfigLoaded())

	// Init new server
	app := server.NewApp(cfg, scfg)

	// Run stream recieving
	app.StreamUC.ServeStreams()

	// Run the server and wait interrupt signal
	if err := app.Run(); err != nil {
		logger.Print(msg.FatalFailureOnServerRunning(err))
	}
}
