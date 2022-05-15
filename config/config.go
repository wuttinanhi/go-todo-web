package config

import "os"

var config *Config

// config struct
type Config struct {
	MYSQL_HOST     string
	MYSQL_DATABASE string
	MYSQL_USER     string
	MYSQL_PASSWORD string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// get config singleton
func GetConfig() Config {
	if config == nil {
		config = &Config{
			MYSQL_HOST:     getEnv("MYSQL_HOST", "127.0.0.1:3306"),
			MYSQL_DATABASE: getEnv("MYSQL_DATABASE", "gotodoweb"),
			MYSQL_USER:     getEnv("MYSQL_USER", "root"),
			MYSQL_PASSWORD: getEnv("MYSQL_PASSWORD", "password"),
		}
	}

	return *config
}
