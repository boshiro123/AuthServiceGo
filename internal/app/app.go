package app

import (
	grpcapp "auth-service-go/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	gprcPort int,
	tokenTTL time.Duration,
) *App {
	grpcApp := grpcapp.New(log, gprcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
