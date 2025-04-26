package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"person-enrichment-api/config"
	"person-enrichment-api/internal/api"
	"person-enrichment-api/internal/migrator"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DATABASE.Username, cfg.DATABASE.Password, cfg.DATABASE.Host, cfg.DATABASE.Port, cfg.DATABASE.DbName)
	ctx := context.Background()
	log := setupLogger(cfg.Env)

	server := api.NewAPIServer(cfg.HTTPServer.Address, log)

	migrator.Migrate(log, dsn, cfg.DATABASE.MigrationsPath)

	done := make(chan bool)
	go func() {
		if err := server.Run(ctx, cfg); err != nil {
			slog.Error("server error", slog.String("err", err.Error()))
		}
		done <- true
	}()

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", slog.String("err", err.Error()))
	}
	<-done
	slog.Info("Server exiting")
}

const (
	envDev  = "dev"
	envProd = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}))
	}
	return log
}
