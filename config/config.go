package config

import (
	"os"
)

type Config struct {
	Port           string
	BackendAPIURL  string
	BackendTimeout int
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		BackendAPIURL:  getEnv("BACKEND_API_URL", "http://localhost:5007"),
		BackendTimeout: 30, // seconds
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
