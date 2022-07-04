package config

import "sync"

type Config struct {
	DBConnectionLatencyMilliseconds int
	DBConnectionShowStatus          bool
	DBConnectionTimeoutSeconds      int
	DBDriver                        string
	DBHost                          string
	DBName                          string
	DBPort                          int
	DBSSLEnable                     bool
	DBUsername                      string

	DBPassword string

	DBOConnectionLatencyMilliseconds int
	DBOConnectionShowStatus          bool
	DBOConnectionTimeoutSeconds      int
	DBODriver                        string
	DBOHost                          string
	DBOName                          string
	DBOPort                          int
	DBOSSLEnable                     bool
	DBOUsername                      string

	DBOPassword string

	HashingPasswordSalt    string
	HashingTokenSigningKey string

	PaginationGetLimitDefault int

	ServerDebugEnable         bool
	ServerMaxHeaderBytes      int
	ServerHost                string
	ServerPort                int
	ServerReadTimeoutSeconds  int
	ServerWriteTimeoutSeconds int

	SessionTTLHours int

	StreamICEServersMutex            sync.RWMutex
	StreamICEServers                 []string
	StreamLink                       string
	StreamSnapshotPeriodSeconds      int
	StreamSnapshotShowStatus         bool
	StreamSnapshotsEnable            bool
	StreamStreamsUpdatePeriodSeconds int

	ServerIP             string
	IsVideoConcatinating bool
}
