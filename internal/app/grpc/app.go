package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "auth-service-go/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log  *slog.Logger
	gRPC *grpc.Server
	port int
}

func New(
	log *slog.Logger,
	authService authgrpc.Auth,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.RegisterServerAPI(gRPCServer, authService)

	return &App{
		log:  log,
		gRPC: gRPCServer,
		port: port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
		slog.String("file", "internal/app/grpc/app.go"),
	)

	log.Info("starting grpc server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server running", slog.String("address", l.Addr().String()),
		slog.String("file", "internal/app/grpc/app.go"))

	if err := a.gRPC.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping grpc server", slog.Int("port", a.port))

	a.gRPC.GracefulStop()
}
