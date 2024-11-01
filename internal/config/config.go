package config

type Config struct {
	Host string
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "8080",
	}
}
