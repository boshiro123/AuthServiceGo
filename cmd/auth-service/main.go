package main

import (
	"auth-service-go/internal/app"
	"auth-service-go/pkg/config"
	"auth-service-go/pkg/logger"
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("its start working")

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting auth service",
		slog.Any("config", cfg),
	)

	application := app.New(log, cfg.GRPC.Port, cfg.TokenTTL)
	// TODO: initialize app
	application.GRPCServer.MustRun()
	// TODO: run gr
}
