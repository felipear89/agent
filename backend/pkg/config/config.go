package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log/slog"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string      `env:"PORT"`
	ResponseTimeout string      `env:"TIMEOUT" envDefault:"30s"`
	DatabaseURL     string      `env:"DATABASE_URL"`
	AuthConfig      *AuthConfig `env:",init"`
}

func (c Config) TimeoutDuration() time.Duration {
	duration, err := time.ParseDuration(c.ResponseTimeout)
	if err != nil {
		slog.Error("Invalid TIMEOUT format, using default 30s", "error", err)
		duration = 30 * time.Second
	}
	return duration
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, using system environment variables")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	if err := cfg.AuthConfig.LoadJWT(); err != nil {
		return nil, err
	}

	slog.Info("DatabaseURL loaded", "url", cfg.DatabaseURL)

	return &cfg, nil
}
