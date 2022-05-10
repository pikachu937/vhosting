package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/server"
)

func main() {
	// Load environment file
	if err := godotenv.Load("./configs/.env"); err != nil {
		logger.Print(msg.FatalFailedToLoadEnvironmentFile(err))
		return
	}
	logger.Print(msg.InfoEnvironmentVarsLoaded())

	var cfg config_tool.Config
	var err error

	// Load config file
	if cfg, err = config_tool.LoadConfig("./configs/config.yml"); err != nil {
		logger.Print(msg.FatalFailedToLoadConfigFile(err))
		return
	}
	logger.Print(msg.InfoConfigVarsLoaded())

	app := server.NewApp(cfg)

	// Run the server and wait for Ctrl+C combination from keyboard to shut down
	if err := app.Run(); err != nil {
		logger.Print(msg.FatalFailureOnServerRunning(err))
	}
}
