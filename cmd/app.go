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
	handler "github.com/mikerumy/vhservice/pkg/handler/userinterface"
	"github.com/mikerumy/vhservice/pkg/logger"
	repository "github.com/mikerumy/vhservice/pkg/repository/userinterface"
	service "github.com/mikerumy/vhservice/pkg/service/userinterface"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Setting up logger
	logger.InitLogger()

	// Reading config file
	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to initializing config: %s\n", err.Error())
	}

	// Reading covert parameters from .env file e.g. DB Password, ...
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to loading env variables: %s\n", err.Error())
	}

	/* ======== Applying config settings ======= */
	svCfg := vhs.SVConfig{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}

	dbCfg := vhs.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	repos := repository.NewRepository(dbCfg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := gin.New()
	userInterface := router.Group("/userinterface")
	{
		userInterface.POST("/", handlers.POSTUser)
		userInterface.GET("/:id", handlers.GETUser)
		userInterface.GET("/all", handlers.GETAllUsers)
		userInterface.PUT("/:id", handlers.PUTUser)
		userInterface.PATCH("/:id", handlers.PATCHUser)
		userInterface.DELETE("/:id", handlers.DELETEUser)
	}
	/* ========================================= */

	// Starting Server
	srv := new(vhs.Server)
	go func() {
		srv.Run(svCfg, router)
	}()

	// Shutting Down Server
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
