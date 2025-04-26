package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"person-enrichment-api/config"
	"person-enrichment-api/internal/api"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	dsn := ""
	server := api.NewAPIServer(cfg.HTTPServer.Address, dsn)
	ctx := context.Background()

	done := make(chan bool)
	go func() {
		if err := server.Run(cfg, ctx); err != nil {
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
