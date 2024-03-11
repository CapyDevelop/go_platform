package logger

import (
	"log"
	"log/slog"
	"os"
)

func NewSlogLogger(level, encoding string) *slog.Logger {
	var (
		logger *slog.Logger
		opt    *slog.HandlerOptions
	)

	switch level {
	case "DEBUG":
		opt = &slog.HandlerOptions{Level: slog.LevelDebug}
	case "INFO":
		opt = &slog.HandlerOptions{Level: slog.LevelInfo}
	case "WARN":
		opt = &slog.HandlerOptions{Level: slog.LevelWarn}
	case "ERROR":
		opt = &slog.HandlerOptions{Level: slog.LevelError}
	default:
		log.Panicf("wrong log level: %v", level)
	}

	switch encoding {
	case "console":
		logger = slog.New(slog.NewTextHandler(os.Stdout, opt))
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opt))
	default:
		log.Panicf("wrong log encoding: %v", encoding)
	}

	return logger
}
