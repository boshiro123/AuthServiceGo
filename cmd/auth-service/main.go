package main

import (
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

	// TODO: initialize app

	// TODO: run gr
}
