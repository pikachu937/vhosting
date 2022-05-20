package config

import (
	"os"
	"strconv"
	"strings"

	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (Config, error) {
	var err error

	// Parse config file path
	path = path[:len(path)-4]
	lastDirIndex := strings.LastIndex(path, "/")
	viper.AddConfigPath(path[:lastDirIndex+1])
	viper.SetConfigName(path[lastDirIndex+1:])

	// Load data from config file
	var cfg Config
	if err = viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	var cvarName string

	cvarName = "db.connTimeoutSeconds"
	dbConnTimeoutSeconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		dbConnTimeoutSeconds = 5
		logger.Print(msg.WarningCannotConvertCvar(cvarName, dbConnTimeoutSeconds, err))
	}

	cvarName = "db.connLatencyMilliseconds"
	dbConnLatencyMilliseconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		dbConnLatencyMilliseconds = 100
		logger.Print(msg.WarningCannotConvertCvar(cvarName, dbConnLatencyMilliseconds, err))
	}

	cvarName = "db.logConnStatus"
	dbLogConnStatus, err := strconv.ParseBool(viper.GetString(cvarName))
	if err != nil {
		dbLogConnStatus = true
		logger.Print(msg.WarningCannotConvertCvar(cvarName, dbLogConnStatus, err))
	}

	cvarName = "server.debugMode"
	serverDebugMode, err := strconv.ParseBool(viper.GetString(cvarName))
	if err != nil {
		serverDebugMode = true
		logger.Print(msg.WarningCannotConvertCvar(cvarName, serverDebugMode, err))
	}

	cvarName = "server.maxHeaderBytes"
	serverMaxHeaderBytes, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		serverMaxHeaderBytes = 1048576 // 1 megabyte
		logger.Print(msg.WarningCannotConvertCvar(cvarName, serverMaxHeaderBytes, err))
	}

	cvarName = "server.readTimeoutSeconds"
	serverReadTimeoutSeconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		serverReadTimeoutSeconds = 15
		logger.Print(msg.WarningCannotConvertCvar(cvarName, serverReadTimeoutSeconds, err))
	}

	cvarName = "server.writeTimeoutSeconds"
	serverWriteTimeoutSeconds, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		serverWriteTimeoutSeconds = 15
		logger.Print(msg.WarningCannotConvertCvar(cvarName, serverWriteTimeoutSeconds, err))
	}

	cvarName = "session.ttlHours"
	sessionTTLHours, err := strconv.Atoi(viper.GetString(cvarName))
	if err != nil {
		sessionTTLHours = 336 // 2 weeks
		logger.Print(msg.WarningCannotConvertCvar(cvarName, sessionTTLHours, err))
	}

	cfg = Config{
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
