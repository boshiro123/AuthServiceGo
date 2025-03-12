package main

import (
	"auth-service-go/internal/app"
	"auth-service-go/internal/repository"
	"auth-service-go/internal/services"
	"auth-service-go/pkg/config"
	"auth-service-go/pkg/logger"
	"auth-service-go/pkg/store"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	sqlc "auth-service-go/internal/db/sqlc"
)

func main() {
	fmt.Println("its start working")
	//load config
	cfg := config.MustLoad()

	fmt.Println(cfg)
	//setup logger
	log := logger.SetupLogger(cfg.Env)

	log.Info("starting auth service",
		slog.Any("config", cfg),
		slog.String("file", "cmd/auth-service/main.go"),
	)
	// TODO: run gr

	//run postgres

	db, err := store.NewPostgresClient()
	if err != nil {
		log.Error("failed to connect to postgres", slog.String("error", err.Error()))
		os.Exit(1)
	}

	dbQ := sqlc.New(db)

	defer func() {
		store.ClosePool(db)
	}()

	authRepo := repository.NewAuth(db, dbQ)

	authService := services.NewAuthService(authRepo)

	//setup application
	application := app.New(log, cfg.GRPC.Port, *authService, cfg.TokenTTL)
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGINT)
	sign := <-stop

	log.Info("stopping auth service", slog.String("signal", sign.String()))
	application.GRPCServer.Stop()
	log.Info("auth service stopped")
}
