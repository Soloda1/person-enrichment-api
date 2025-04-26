package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"person-enrichment-api/config"
	"person-enrichment-api/internal/handlers/ping"
	"person-enrichment-api/internal/middleware/requestLogger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type APIServer struct {
	address string
	dbUrl   string
	router  *gin.Engine
	server  *http.Server
}

func NewAPIServer(address string, dbUrl string) *APIServer {
	return &APIServer{address: address, dbUrl: dbUrl}
}

func (s *APIServer) Run(cfg *config.Config, ctx context.Context) error {
	log := setupLogger(cfg.Env)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(requestLogger.RequestLoggerMiddleware(log))

	router.GET("/ping", ping.Ping)

	s.server = &http.Server{
		Addr:         s.address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	log.Info("Starting server", slog.String("address", s.address))
	log.Debug("Debug logger enabled")

	return s.server.ListenAndServe()

}

func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

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
