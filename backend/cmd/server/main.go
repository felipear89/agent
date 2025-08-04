package main

import (
	"context"
	"github.com/felipear89/agent/pkg/auth"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/felipear89/agent/pkg/config"
	"github.com/felipear89/agent/pkg/server"
	"github.com/felipear89/agent/pkg/user"
)

func main() {
	// config.SetDefaultLogger()

	slog.Info("Starting application")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	userSvc := user.Register()
	authSvc := auth.Register(cfg)

	srv := server.New(&server.Config{
		Port:         cfg.ServerPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
		Timeout:      cfg.TimeoutDuration(),
		BasePath:     "/",
	})

	api := srv.RegisterAPIRoutes()
	auth.NewHandler(api, authSvc, userSvc)

	api.Use(authSvc.AuthMiddleware())
	user.NewHandler(api, userSvc)

	// Start the server in a goroutine to allow graceful shutdown
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- srv.Start()
	}()

	// Setup graceful shutdown
	gracefulShutdown(serverErrors, cfg, srv)
}

func gracefulShutdown(serverErrors chan error, cfg *config.Config, srv *server.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		slog.Error("Error starting server", "error", err)
		os.Exit(1)

	case <-stop:
		slog.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), cfg.TimeoutDuration())
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Server forced to shutdown:", "error", err)
			os.Exit(1)
		}

		slog.Info("Application stopped")
	}
}
