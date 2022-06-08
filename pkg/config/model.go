package config

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

	StreamSnapshotPeriodSeconds int
	StreamSnapshotShowStatus    bool
	StreamSnapshotsEnable       bool
}
