package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/pkg/handler"
	"github.com/mikerumy/vhosting/pkg/service"
	"github.com/mikerumy/vhosting/pkg/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Set up logger
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Set up reader of config file
	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to initializing config: %s\n", err.Error())
	}

	// Set up reader of environment variables from .env file (DB Password)
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to loading env variables: %s\n", err.Error())
	}

	// Read server part of config settings
	svCfg := vh.SVConfig{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}

	// Read DB part of config settings
	dbCfg := vh.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	// Apply DB part of config
	stors := storage.NewStorage(dbCfg)
	services := service.NewService(stors)
	handlers := handler.NewHandler(services)

	// Init Routes
	router := gin.New()
	userInterface := router.Group("/user-interface")
	{
		userInterface.POST("/", handlers.POSTUser)
		userInterface.GET("/:id", handlers.GETUser)
		userInterface.GET("/all", handlers.GETAllUsers)
		userInterface.PATCH("/:id", handlers.PATCHUser)
		userInterface.DELETE("/:id", handlers.DELETEUser)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp)
		auth.POST("/sign-in", handlers.SignIn)
	}

	// Start Server and init server part of config
	srv := new(vh.Server)
	go func() {
		srv.Run(svCfg, router)
	}()

	// Shut Down Server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	time.Sleep(2 * time.Second)

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
