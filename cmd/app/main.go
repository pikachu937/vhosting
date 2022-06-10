package main

import (
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/config"
	"github.com/mikerumy/vhosting/pkg/config_stream"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/server"
)

func main() {
	// Create log storage
	logs := make([]*logger.Log, 0, 100)

	// Load environments
	if err := godotenv.Load("./configs/.env"); err != nil {
		logs = append(logs, msg.FatalFailedToLoadEnvironmentFile(err))
		return
	}
	logs = append(logs, msg.InfoEnvironmentsLoaded())

	// Load config
	cfg, err := config.LoadConfig("./configs/config.yml")
	if err != nil {
		logs = append(logs, msg.FatalFailedToLoadConfigFile(err))
		return
	}
	logs = append(logs, msg.InfoConfigLoaded())

	// ==============================
	logger.Print(logs[0])
	logger.Print(logs[1])
	// ==============================

	// Load stream config
	scfg, err := config_stream.LoadConfig("./configs/stream_config.json")
	if err != nil {
		logger.Print(msg.FatalFailedToLoadStreamConfigFile(err))
		return
	}
	scfg.Server.HTTPPort = strconv.Itoa(cfg.ServerPort)
	logger.Print(msg.InfoStreamConfigLoaded())

	// Init new server
	app := server.NewApp(cfg, scfg)

	// Run stream recieving
	go app.StreamUC.ServeStreams()

	// Run the server and wait interrupt signal
	if err := app.Run(); err != nil {
		logger.Print(msg.FatalFailureOnServerRunning(err))
	}
}
