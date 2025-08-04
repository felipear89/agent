package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     *Config
	httpServer *http.Server
	router     *gin.Engine
}

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Timeout      time.Duration
	BasePath     string
}

func New(cfg *Config) *Server {
	router := gin.New()

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	server := &Server{
		httpServer: srv,
		router:     router,
		config:     cfg,
	}

	server.setupMiddleware()

	return server
}

func (s *Server) Start() error {
	slog.Info("Server starting", "address", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) HttpServer() *http.Server {
	return s.httpServer
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
