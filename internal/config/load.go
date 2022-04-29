package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
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

	var cvarName string

	cvarName = "db.connTimeoutSeconds"
	dbConnTimeoutSeconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		dbConnTimeoutSeconds = 5
		response.WarningCannotConvertCvar(cvarName, dbConnTimeoutSeconds, err)
	}

	cvarName = "db.connLatencyMilliseconds"
	dbConnLatencyMilliseconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		dbConnLatencyMilliseconds = 100
		response.WarningCannotConvertCvar(cvarName, dbConnLatencyMilliseconds, err)
	}

	cvarName = "db.logConnStatus"
	dbLogConnStatus, err := strconv.ParseBool(viper.GetString(cvarName))
	if err != nil {
		dbLogConnStatus = true
		response.WarningCannotConvertCvar(cvarName, dbLogConnStatus, err)
	}

	cvarName = "server.debugMode"
	serverDebugMode, err := strconv.ParseBool(viper.GetString(cvarName))
	if err != nil {
		serverDebugMode = true
		response.WarningCannotConvertCvar(cvarName, serverDebugMode, err)
	}

	cvarName = "server.maxHeaderBytes"
	serverMaxHeaderBytes, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		serverMaxHeaderBytes = 1048576 // 1 megabyte
		response.WarningCannotConvertCvar(cvarName, serverMaxHeaderBytes, err)
	}

	cvarName = "server.readTimeoutSeconds"
	serverReadTimeoutSeconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		serverReadTimeoutSeconds = 15
		response.WarningCannotConvertCvar(cvarName, serverReadTimeoutSeconds, err)
	}

	cvarName = "server.writeTimeoutSeconds"
	serverWriteTimeoutSeconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		serverWriteTimeoutSeconds = 15
		response.WarningCannotConvertCvar(cvarName, serverWriteTimeoutSeconds, err)
	}

	cvarName = "session.ttlHours"
	sessionTTLHours, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		sessionTTLHours = 336 // 2 weeks
		response.WarningCannotConvertCvar(cvarName, sessionTTLHours, err)
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
