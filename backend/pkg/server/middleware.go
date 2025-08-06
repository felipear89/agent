package server

import (
	"github.com/felipear89/agent/pkg/server/middleware"
)

func (s *Server) setupMiddleware() {
	s.router.Use(middleware.DefaultTimeout(s.config.Timeout))
	s.router.Use(middleware.Recovery())
	s.router.Use(middleware.Logger())
}
