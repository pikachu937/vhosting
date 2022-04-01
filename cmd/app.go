package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	vhs "github.com/mikerumy/vhservice"
	handler "github.com/mikerumy/vhservice/pkg/handler/userinterface"
	repositorydb "github.com/mikerumy/vhservice/pkg/repository/db"
	repository "github.com/mikerumy/vhservice/pkg/repository/userinterface"
	service "github.com/mikerumy/vhservice/pkg/service/userinterface"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var conf repositorydb.Config

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to loading env variables: %s", err.Error())
	}

	conf = repositorydb.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db := repositorydb.NewPostgresConnection(conf)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(vhs.Server)
	go func() {
		host := viper.GetString("host")
		port := viper.GetString("port")
		if err := srv.Run(host, port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("failed to running server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	logrus.Printf("server shut down\n")

	// if err := db.Close(); err != nil {
	// 	logrus.Errorf("error occured on db connection close: %s", err.Error())
	// }

	// logrus.Printf("closed connection to db\n")
}

func initConfig() error {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
