package app

import (
	grpcapp "github.com/CapyDevelop/go_platform/skeleton/internal/app/grpc"
	"github.com/CapyDevelop/go_platform/skeleton/internal/config"
	"log"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
	log        log.Logger
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	grpcApp := grpcapp.New(log, cfg.GRPC.Port)

	return &App{
		GRPCServer: grpcApp,
	}
}
