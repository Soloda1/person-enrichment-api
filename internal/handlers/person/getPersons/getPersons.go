package getPersons

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"person-enrichment-api/internal/repository/person"
	utils "person-enrichment-api/internal/utils/error"
	"person-enrichment-api/internal/utils/logger"
)

type Response struct {
	Status  string           `json:"status"`
	Error   string           `json:"error,omitempty"`
	Persons []*person.Person `json:"person,omitempty"`
}

type PersonsProvider interface {
	GetAllPersons(ctx context.Context) ([]*person.Person, error)
}

func New(log *logger.Logger, service PersonsProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("GetAllPersons called")

		personModel, err := service.GetAllPersons(c.Request.Context())
		if err != nil {
			log.Debug("failed to get person by id", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, Response{
			Status:  http.StatusText(http.StatusOK),
			Persons: personModel,
		})
	}
}
