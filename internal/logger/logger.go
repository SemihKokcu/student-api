package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(format string, level string) *slog.Logger {
	var handler slog.Handler

	// Log seviyesini belirle
	var slogLevel slog.Level
	switch level {
	case "debug":
		slogLevel = slog.LevelDebug
	case "error":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: slogLevel}

	// Formata karar ver
	if format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
}
