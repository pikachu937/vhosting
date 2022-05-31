package config

type Config struct {
	DBConnTimeoutSeconds      int
	DBConnLatencyMilliseconds int
	DBDriver                  string
	DBHost                    string
	DBLogConnStatus           bool
	DBName                    string
	DBPassword                string
	DBPort                    string
	DBSSLMode                 string
	DBUsername                string

	HashingPasswordSalt    string
	HashingTokenSigningKey string

	ServerDebugMode           bool
	ServerMaxHeaderBytes      int
	ServerPort                string
	ServerReadTimeoutSeconds  int
	ServerWriteTimeoutSeconds int

	SessionTTLHours int

	ServerIP string
}
