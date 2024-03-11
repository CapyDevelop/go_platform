package servergrpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"

	hellogrpc "github.com/CapyDevelop/go_platform/skeleton/internal/controllers/grpc/hello"
	"log/slog"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
	host       string
}

func New(
	log *slog.Logger,
	port int,
) *App {
	server := grpc.NewServer()

	hellogrpc.Register(server)

	return &App{
		log:        log,
		gRPCServer: server,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "servergrpc.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.port))
	if err != nil {
		return fmt.Errorf("failed on net.Listen: %w", err)
	}

	log.Info("grpc server is running", slog.String("addr", lis.Addr().String()))

	if err := a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("failed on gRPCServer.Serve: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "servergrpc.Run"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
