package config

import (
	"os"
	"time"
)

type Config struct {
	ServerPort      string
	ResponseTimeout int
}

func (c Config) TimeoutDuration() time.Duration {
	return time.Duration(c.ResponseTimeout) * time.Second
}

func LoadConfig() *Config {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		ServerPort:      port,
		ResponseTimeout: 5,
	}
}
