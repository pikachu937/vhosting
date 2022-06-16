package db_connect

type DBConfig struct {
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
}
