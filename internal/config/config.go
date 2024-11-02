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
		DBName:         "mydb",
		DBUser:         "name",
		DBPassword:     "password",
		DBDisableTLS:   false,
		DBMaxOpenConns: 10,
		DBHost:         "localhost",
		DBPort:         "5432",
	}
}
