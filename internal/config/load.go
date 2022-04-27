package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/mikerumy/vhosting2/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig(configPath string) (models.Config, error) {
	// Parse config file path
	configPath = configPath[:len(configPath)-4]
	var lastDirIndex int = strings.LastIndex(configPath, "/")
	viper.AddConfigPath(configPath[:lastDirIndex+1])
	viper.SetConfigName(configPath[lastDirIndex+1:])

	// Load data from config file
	var cfg models.Config
	var err error
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	var varName string
	const message = "Cannot convert cvar \"%s\". Set default value: %d. Error: %s.\n"

	varName = "db.connTimeoutSeconds"
	dbConnTimeoutSeconds, err := strconv.Atoi(viper.GetString(varName))
	if err != nil {
		dbConnTimeoutSeconds = 5
		logrus.Warningf(message, varName, dbConnTimeoutSeconds, err.Error())
	}

	varName = "db.connLatencyMilliseconds"
	dbConnLatencyMilliseconds, err := strconv.Atoi(viper.GetString(varName))
	if err != nil {
		dbConnLatencyMilliseconds = 100
		logrus.Warningf(message, varName, dbConnLatencyMilliseconds, err.Error())
	}

	varName = "db.logConnStatus"
	dbLogConnStatus, err := strconv.ParseBool(viper.GetString(varName))
	if err != nil {
		dbLogConnStatus = true
		logrus.Warningf(message, varName, dbLogConnStatus, err.Error())
	}

	varName = "server.debugMode"
	serverDebugMode, err := strconv.ParseBool(viper.GetString(varName))
	if err != nil {
		serverDebugMode = true
		logrus.Warningf(message, varName, serverDebugMode, err.Error())
	}

	varName = "server.maxHeaderBytes"
	serverMaxHeaderBytes, err := strconv.Atoi(viper.GetString(varName))
	if err != nil {
		serverMaxHeaderBytes = 1048576 // 1 megabyte
		logrus.Warningf(message, varName, serverMaxHeaderBytes, err.Error())
	}

	varName = "server.readTimeoutSeconds"
	serverReadTimeoutSeconds, err := strconv.Atoi(viper.GetString(varName))
	if err != nil {
		serverReadTimeoutSeconds = 15
		logrus.Warningf(message, varName, serverReadTimeoutSeconds, err.Error())
	}

	varName = "server.writeTimeoutSeconds"
	serverWriteTimeoutSeconds, err := strconv.Atoi(viper.GetString(varName))
	if err != nil {
		serverWriteTimeoutSeconds = 15
		logrus.Warningf(message, varName, serverWriteTimeoutSeconds, err.Error())
	}

	varName = "session.ttlHours"
	sessionTTLHours, err := strconv.Atoi(viper.GetString(varName))
	if err != nil {
		sessionTTLHours = 336 // 2 weeks
		logrus.Warningf(message, varName, sessionTTLHours, err.Error())
	}

	cfg = models.Config{
		DBConnTimeoutSeconds:      dbConnTimeoutSeconds,
		DBConnLatencyMilliseconds: dbConnLatencyMilliseconds,
		DBDriver:                  os.Getenv("DB_DRIVER"),
		DBHost:                    viper.GetString("db.host"),
		DBLogConnStatus:           dbLogConnStatus,
		DBName:                    viper.GetString("db.name"),
		DBPassword:                os.Getenv("DB_PASSWORD"),
		DBPort:                    viper.GetString("db.port"),
		DBSSLMode:                 viper.GetString("db.sslmode"),
		DBUsername:                viper.GetString("db.username"),

		HashingPasswordSalt:    os.Getenv("HASHING_PASSWORD_SALT"),
		HashingTokenSigningKey: os.Getenv("HASHING_TOKEN_SIGNING_KEY"),

		ServerDebugMode:           serverDebugMode,
		ServerMaxHeaderBytes:      serverMaxHeaderBytes,
		ServerPort:                viper.GetString("server.port"),
		ServerReadTimeoutSeconds:  serverReadTimeoutSeconds,
		ServerWriteTimeoutSeconds: serverWriteTimeoutSeconds,

		SessionTTLHours: sessionTTLHours,
	}

	return cfg, nil
}
