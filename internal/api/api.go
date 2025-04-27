package api

import (
	"context"
	"log/slog"
	"net/http"
	"person-enrichment-api/config"
	"person-enrichment-api/internal/handlers/person/createPerson"
	"person-enrichment-api/internal/handlers/person/deletePerson"
	"person-enrichment-api/internal/handlers/person/getPerson"
	"person-enrichment-api/internal/handlers/person/getPersons"
	"person-enrichment-api/internal/handlers/person/updatePerson"
	"person-enrichment-api/internal/handlers/ping"
	"person-enrichment-api/internal/middleware/requestLogger"
	"person-enrichment-api/internal/service/person"
	"person-enrichment-api/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	address       string
	log           *logger.Logger
	router        *gin.Engine
	server        *http.Server
	personService person.PersonService
}

func NewAPIServer(address string, log *logger.Logger, personService person.PersonService) *APIServer {
	return &APIServer{
		address:       address,
		log:           log,
		personService: personService}
}

func (s *APIServer) Run(ctx context.Context, cfg *config.Config) error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(requestLogger.RequestLoggerMiddleware(s.log))

	router.GET("/ping", ping.Ping)

	router.POST("/create", createPerson.New(s.log, s.personService))
	router.DELETE("/delete/:id", deletePerson.New(s.log, s.personService))
	router.GET("/person/:id", getPerson.New(s.log, s.personService))
	router.GET("/persons", getPersons.New(s.log, s.personService))
	router.PUT("/update", updatePerson.New(s.log, s.personService))

	s.server = &http.Server{
		Addr:         s.address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	s.log.Info("Starting server", slog.String("address", s.address))
	s.log.Debug("Debug logger enabled")

	return s.server.ListenAndServe()
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
