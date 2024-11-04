package config

type Config struct {
	Host           string
	Port           string
	DBName         string
	DBUser         string
	DBPassword     string
	DBDisableTLS   bool
	DBMaxOpenConns int
	DBHost         string
	DBPort         string
}

func LoadConfig() *Config {
	return &Config{
		Host:           "localhost",
		Port:           "8080",
		DBName:         "blogdb",
		DBUser:         "root",
		DBPassword:     "secret",
		DBDisableTLS:   true,
		DBMaxOpenConns: 10,
		DBHost:         "localhost",
		DBPort:         "5321",
	}
}
