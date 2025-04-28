package getPersons

import (
	"context"
	"log/slog"
	"net/http"
	"person-enrichment-api/config"
	"person-enrichment-api/internal/models"
	utils "person-enrichment-api/internal/utils/error"
	"person-enrichment-api/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string           `json:"status"`
	Error   string           `json:"error,omitempty"`
	Persons []*models.Person `json:"person,omitempty"`
}

type PersonsProvider interface {
	GetAllPersons(ctx context.Context, filter models.PersonFilter) ([]*models.Person, error)
}

func New(log *logger.Logger, service PersonsProvider, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("GetAllPersons called")

		filter := models.PersonFilter{}

		limitStr := c.DefaultQuery("limit", strconv.Itoa(cfg.HTTPServer.Pagination.DefaultLimit))
		pageStr := c.DefaultQuery("page", strconv.Itoa(cfg.HTTPServer.Pagination.DefaultPage))

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			log.Debug("failed to parse limit", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, "Invalid limit")
			return
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			log.Debug("failed to parse page", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, "Invalid page")
			return
		}

		filter.Offset = (page - 1) * limit
		filter.Limit = limit

		personModel, err := service.GetAllPersons(c.Request.Context(), filter)
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
