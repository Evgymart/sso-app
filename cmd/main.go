package main

import (
	"fmt"
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"

	ssov1 "protos/gen/go/sso"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()
	req := &ssov1.RegisterRequest{
		Email:    "user@example.com",
		Password: "password123",
	}

	logger := setupLogger(config.Env)
	logger.Info(fmt.Sprintf("Request: %+v\n", req))
	logger.Info(fmt.Sprintf("Config: %+v\n", config))

	application := app.New(logger, config.GRPC.Port, config.TokenTTL)
	application.GRPCSrv.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
