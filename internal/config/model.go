package config

type Config struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	DBDriver string
	Password string
}
