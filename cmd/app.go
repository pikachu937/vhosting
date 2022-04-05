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
	vhs "github.com/mikerumy/vhservice"
	"github.com/mikerumy/vhservice/pkg/handler"
	"github.com/mikerumy/vhservice/pkg/service"
	"github.com/mikerumy/vhservice/pkg/storage"
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
	svCfg := vhs.SVConfig{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}

	// Read DB part of config settings
	dbCfg := vhs.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	// Apply DB part of config
	stor := storage.NewStorage(dbCfg)
	services := service.NewService(stor)
	handler := handler.NewHandler(services)

	// Init Routes
	router := gin.New()
	userInterface := router.Group("/user-interface")
	{
		userInterface.POST("/", handler.POSTUser)
		userInterface.GET("/:id", handler.GETUser)
		userInterface.GET("/all", handler.GETAllUsers)
		userInterface.PUT("/:id", handler.PUTUser)
		userInterface.PATCH("/:id", handler.PATCHUser)
		userInterface.DELETE("/:id", handler.DELETEUser)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.SignUp)
		auth.POST("/sign-in", handler.SignIn)
	}

	// Start Server and init server part of config
	srv := new(vhs.Server)
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
