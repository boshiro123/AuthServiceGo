package main

import (
	"auth-service-go/internal/app"
	"auth-service-go/pkg/config"
	"auth-service-go/pkg/logger"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("its start working")

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting auth service",
		slog.Any("config", cfg),
		slog.String("file", "cmd/auth-service/main.go"),
)

	application := app.New(log, cfg.GRPC.Port, cfg.TokenTTL)
	// TODO: initialize app
	go application.GRPCServer.MustRun()
	// TODO: run gr

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGINT)
	sign:=<-stop

	log.Info("stopping auth service", slog.String("signal", sign.String()))
	application.GRPCServer.Stop()
	log.Info("auth service stopped")
}
