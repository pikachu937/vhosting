package config

type Config struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	DBDriver string
}
