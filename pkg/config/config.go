package config

import (
	"os"
	"strings"

	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*Config, error) {
	// Parse config file path
	path = path[:len(path)-4]
	lastDirIndex := strings.LastIndex(path, "/")
	viper.AddConfigPath(path[:lastDirIndex+1])
	viper.SetConfigName(path[lastDirIndex+1:])

	// Load data from config file
	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return &cfg, err
	}

	var cvar string

	cvar = "db.connectionLatencyMilliseconds"
	dbConnectionLatencyMilliseconds := viper.GetInt(cvar)
	if dbConnectionLatencyMilliseconds == 0 {
		dbConnectionLatencyMilliseconds = 100
		logger.Print(msg.WarningCannotConvertCvar(cvar, dbConnectionLatencyMilliseconds))
	}

	cvar = "db.connectionTimeoutSeconds"
	dbConnectionTimeoutSeconds := viper.GetInt(cvar)
	if dbConnectionTimeoutSeconds == 0 {
		dbConnectionTimeoutSeconds = 5
		logger.Print(msg.WarningCannotConvertCvar(cvar, dbConnectionTimeoutSeconds))
	}

	cvar = "db.port"
	dbPort := viper.GetInt(cvar)
	if dbPort == 0 {
		dbPort = 3456
		logger.Print(msg.WarningCannotConvertCvar(cvar, dbPort))
	}

	cvar = "dbo.connectionLatencyMilliseconds"
	dboConnectionLatencyMilliseconds := viper.GetInt(cvar)
	if dboConnectionLatencyMilliseconds == 0 {
		dboConnectionLatencyMilliseconds = 100
		logger.Print(msg.WarningCannotConvertCvar(cvar, dboConnectionLatencyMilliseconds))
	}

	cvar = "dbo.connectionTimeoutSeconds"
	dboConnectionTimeoutSeconds := viper.GetInt(cvar)
	if dboConnectionTimeoutSeconds == 0 {
		dboConnectionTimeoutSeconds = 5
		logger.Print(msg.WarningCannotConvertCvar(cvar, dboConnectionTimeoutSeconds))
	}

	cvar = "dbo.port"
	dboPort := viper.GetInt(cvar)
	if dboPort == 0 {
		dboPort = 3456
		logger.Print(msg.WarningCannotConvertCvar(cvar, dboPort))
	}

	cvar = "pagination.getLimitDefault"
	paginationGetLimitDefault := viper.GetInt(cvar)
	if paginationGetLimitDefault == 0 {
		paginationGetLimitDefault = 30
		logger.Print(msg.WarningCannotConvertCvar(cvar, paginationGetLimitDefault))
	}

	cvar = "server.maxHeaderBytes"
	serverMaxHeaderBytes := viper.GetInt(cvar)
	if serverMaxHeaderBytes == 0 {
		serverMaxHeaderBytes = 1048576 // 1 megabyte
		logger.Print(msg.WarningCannotConvertCvar(cvar, serverMaxHeaderBytes))
	}

	cvar = "server.port"
	serverPort := viper.GetInt(cvar)
	if serverPort == 0 {
		serverPort = 8000
		logger.Print(msg.WarningCannotConvertCvar(cvar, serverPort))
	}

	cvar = "server.readTimeoutSeconds"
	serverReadTimeoutSeconds := viper.GetInt(cvar)
	if serverReadTimeoutSeconds == 0 {
		serverReadTimeoutSeconds = 15
		logger.Print(msg.WarningCannotConvertCvar(cvar, serverReadTimeoutSeconds))
	}

	cvar = "server.writeTimeoutSeconds"
	serverWriteTimeoutSeconds := viper.GetInt(cvar)
	if serverWriteTimeoutSeconds == 0 {
		serverWriteTimeoutSeconds = 15
		logger.Print(msg.WarningCannotConvertCvar(cvar, serverWriteTimeoutSeconds))
	}

	cvar = "session.ttlHours"
	sessionTTLHours := viper.GetInt(cvar)
	if sessionTTLHours == 0 {
		sessionTTLHours = 168 // 7 days
		logger.Print(msg.WarningCannotConvertCvar(cvar, sessionTTLHours))
	}

	cvar = "stream.snapshotPeriodSeconds"
	streamSnapshotPeriodSeconds := viper.GetInt(cvar)
	if streamSnapshotPeriodSeconds == 0 {
		streamSnapshotPeriodSeconds = 60
		logger.Print(msg.WarningCannotConvertCvar(cvar, streamSnapshotPeriodSeconds))
	}

	cfg = Config{
		DBConnectionLatencyMilliseconds: dbConnectionLatencyMilliseconds,
		DBConnectionShowStatus:          viper.GetBool("db.connectionShowStatus"),
		DBConnectionTimeoutSeconds:      dbConnectionTimeoutSeconds,
		DBHost:                          viper.GetString("db.host"),
		DBName:                          viper.GetString("db.name"),
		DBPort:                          dbPort,
		DBSSLEnable:                     viper.GetBool("db.sslEnable"),
		DBUsername:                      viper.GetString("db.username"),
		DBDriver:                        os.Getenv("DB_DRIVER"),
		DBPassword:                      os.Getenv("DB_PASSWORD"),

		DBOConnectionLatencyMilliseconds: dboConnectionLatencyMilliseconds,
		DBOConnectionShowStatus:          viper.GetBool("dbo.connectionShowStatus"),
		DBOConnectionTimeoutSeconds:      dboConnectionTimeoutSeconds,
		DBOHost:                          viper.GetString("dbo.host"),
		DBOName:                          viper.GetString("dbo.name"),
		DBOPort:                          dboPort,
		DBOSSLEnable:                     viper.GetBool("dbo.sslEnable"),
		DBOUsername:                      viper.GetString("dbo.username"),
		DBODriver:                        os.Getenv("DBO_DRIVER"),
		DBOPassword:                      os.Getenv("DBO_PASSWORD"),

		HashingPasswordSalt:    os.Getenv("HASHING_PASSWORD_SALT"),
		HashingTokenSigningKey: os.Getenv("HASHING_TOKEN_SIGNING_KEY"),

		PaginationGetLimitDefault: paginationGetLimitDefault,

		ServerDebugEnable:         viper.GetBool("server.debugEnable"),
		ServerMaxHeaderBytes:      serverMaxHeaderBytes,
		ServerPort:                serverPort,
		ServerReadTimeoutSeconds:  serverReadTimeoutSeconds,
		ServerWriteTimeoutSeconds: serverWriteTimeoutSeconds,

		SessionTTLHours: sessionTTLHours,

		StreamSnapshotPeriodSeconds: streamSnapshotPeriodSeconds,
		StreamSnapshotShowStatus:    viper.GetBool("stream.snapshotShowStatus"),
		StreamSnapshotsEnable:       viper.GetBool("stream.snapshotsEnable"),
	}

	return &cfg, nil
}
