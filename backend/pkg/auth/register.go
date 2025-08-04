package auth

import (
	"github.com/felipear89/agent/pkg/config"
)

func Register(cfg *config.Config) *Service {
	authService := newService(
		cfg.AuthConfig,
	)
	return authService
}
