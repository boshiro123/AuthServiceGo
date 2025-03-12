package app

import (
	grpcapp "auth-service-go/internal/app/grpc"
	"auth-service-go/internal/services"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	port int,
	authService services.AuthService,
	tokenTTL time.Duration,
) *App {

	grpcApp := grpcapp.New(log, &authService, port)

	return &App{
		GRPCServer: grpcApp,
	}
}
