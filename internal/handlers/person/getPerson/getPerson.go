package getPerson

import (
	"context"
	"log/slog"
	"net/http"
	"person-enrichment-api/internal/models"
	utils "person-enrichment-api/internal/utils/error"
	"person-enrichment-api/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type Request struct {
	PersonID int `uri:"id" binding:"required"`
}

type Response struct {
	Status string         `json:"status"`
	Error  string         `json:"error,omitempty"`
	Person *models.Person `json:"person,omitempty"`
}

type PersonByIdProvider interface {
	GetPersonByID(ctx context.Context, personId int) (*models.Person, error)
}

func New(log *logger.Logger, service PersonByIdProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("GetPersonByID called")

		var req Request
		if err := c.ShouldBindUri(&req); err != nil {
			log.Debug("personId is invalid", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusBadRequest, "invalid person id")
			return
		}

		personModel, err := service.GetPersonByID(c.Request.Context(), req.PersonID)
		if err != nil {
			log.Debug("failed to get person by id", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, Response{
			Status: http.StatusText(http.StatusOK),
			Person: personModel,
		})
	}
}
