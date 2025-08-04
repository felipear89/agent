package config

import (
	"log/slog"
	"os"
)

func SetDefaultLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
