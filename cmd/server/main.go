package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"person-enrichment-api/config"
	"person-enrichment-api/external/enrichment"
	"person-enrichment-api/internal/api"
	"person-enrichment-api/internal/migrator"
	"person-enrichment-api/internal/repository"
	personrepository "person-enrichment-api/internal/repository/person"
	personservice "person-enrichment-api/internal/service/person"
	"person-enrichment-api/internal/utils/logger"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DATABASE.Username,
		cfg.DATABASE.Password,
		cfg.DATABASE.Host,
		cfg.DATABASE.Port,
		cfg.DATABASE.DbName)
	ctx := context.Background()
	log := logger.New(cfg.Env)

	if err := migrator.Migrate(log, dsn, cfg.DATABASE.MigrationsPath); err != nil {
		log.Debug("Failed to run migrations", slog.String("error", err.Error()))
		os.Exit(1)
	}

	storage, err := repository.NewStorage(ctx, dsn)
	if err != nil {
		log.Debug("Failed to create storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	enrichmentService := enrichment.NewEnrichmentService()

	personRepo := personrepository.NewRepository(storage, log)
	personService := personservice.NewService(personRepo, log, enrichmentService)

	server := api.NewAPIServer(
		cfg.HTTPServer.Address,
		log,
		personService)

	done := make(chan bool)
	go func() {
		if err := server.Run(ctx, cfg); err != nil {
			log.Debug("server error", slog.String("error", err.Error()))
		}
		done <- true
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("server shutdown error", slog.String("error", err.Error()))
	}
	storage.Close()
	<-done
	log.Info("Server exiting")
}
