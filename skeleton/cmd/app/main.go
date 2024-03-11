package main

import (
	"context"
	"fmt"
	"github.com/CapyDevelop/go_platform/skeleton/internal/app"
	"github.com/CapyDevelop/go_platform/skeleton/internal/config"
	"github.com/CapyDevelop/go_platform/skeleton/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()

	fmt.Printf("%#v", cfg)

	slogLogger := logger.NewSlogLogger(cfg.Logger.Level, cfg.Logger.Encoding)

	application := app.New(slogLogger, cfg)

	go func() {
		application.GRPCServer.MustRun()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)

	defer func() {
		signal.Stop(sigCh)
	}()

	sig := <-sigCh

	cancel()
	slogLogger.Info("received signal, exiting", "signal", sig)
}
