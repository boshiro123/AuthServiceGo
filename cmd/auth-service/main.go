package main

import (
	"auth-service-go/internal/handler"
	"auth-service-go/internal/repository"
	"auth-service-go/internal/services"
	"auth-service-go/pkg/config"
	"auth-service-go/pkg/logger"
	"auth-service-go/pkg/store"
	"log/slog"

	sqlc "auth-service-go/internal/db/sqlc"
)

func main() {
	//Load config
	cfg := config.MustLoad()

	//setup logger
	log := logger.SetupLogger(cfg.Env)

	log.Info("starting auth service",
		slog.Any("config", cfg),
		slog.String("file", "cmd/auth-service/main.go"),
	)
	// TODO: run grpc server

	//run postgres

	db, err := store.NewPostgresClient()
	if err != nil {
		log.Error("failed to connect to postgres", slog.String("error", err.Error()))
		return
	}

	dbQ := sqlc.New(db)
	defer store.ClosePool(db)

	authRepo := repository.NewAuth(db, dbQ)

	authService := services.NewAuthService(authRepo)

	handlers := handler.NewHandler(authService, log)

	handlers.InitRoutes()

}
