package main

import (
	"context"
	"fmt"
	"github.com/CapyDevelop/go_platform/skeleton/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	logger := NewSlogLogger(cfg.Logger.Level, cfg.Logger.Encoding)

	fmt.Println(ctx, cfg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)

	defer func() {
		signal.Stop(sigCh)
	}()

	sig := <-sigCh

	cancel()

	logger.Info("received signal, exiting", "signal", sig)
}

func NewSlogLogger(level, encoding string) *slog.Logger {
	var (
		log *slog.Logger
		opt *slog.HandlerOptions
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
	}

	switch encoding {
	case "console":
		log = slog.New(slog.NewTextHandler(os.Stdout, opt))
	case "json":
		log = slog.New(slog.NewJSONHandler(os.Stdout, opt))
	}

	return log
}
