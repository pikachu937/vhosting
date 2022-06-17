package config

import "sync"

type Config struct {
	DBConnectionLatencyMilliseconds int
	DBConnectionShowStatus          bool
	DBConnectionTimeoutSeconds      int
	DBHost                          string
	DBName                          string
	DBPort                          int
	DBSSLEnable                     bool
	DBUsername                      string
	DBDriver                        string
	DBPassword                      string

	DBOConnectionLatencyMilliseconds int
	DBOConnectionShowStatus          bool
	DBOConnectionTimeoutSeconds      int
	DBOHost                          string
	DBOName                          string
	DBOPort                          int
	DBOSSLEnable                     bool
	DBOUsername                      string
	DBODriver                        string
	DBOPassword                      string

	HashingPasswordSalt    string
	HashingTokenSigningKey string

	PaginationGetLimitDefault int

	ServerDebugEnable         bool
	ServerMaxHeaderBytes      int
	ServerPort                int
	ServerReadTimeoutSeconds  int
	ServerWriteTimeoutSeconds int

	SessionTTLHours int

	ServerIP string

	StreamICEServersMutex       sync.RWMutex
	StreamICEServers            []string
	StreamLink                  string
	StreamSnapshotPeriodSeconds int
	StreamSnapshotShowStatus    bool
	StreamSnapshotsEnable       bool
}
