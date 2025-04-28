package createPerson

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
	Name       string  `json:"name" binding:"required"`
	Surname    string  `json:"surname" binding:"required"`
	Patronymic *string `json:"patronymic" `
}

type Response struct {
	Status string         `json:"status"`
	Error  string         `json:"error,omitempty"`
	Person *models.Person `json:"person,omitempty"`
}

type PersonCreator interface {
	CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
}

func New(log *logger.Logger, service PersonCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("CreatePerson called")

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Debug("failed to bind request body", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusBadRequest, "invalid request body")
			return
		}

		personModel := &models.Person{
			Name:       req.Name,
			Surname:    req.Surname,
			Patronymic: req.Patronymic,
		}

		createdPerson, err := service.CreatePerson(c.Request.Context(), personModel)
		if err != nil {
			log.Debug("failed to create person", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, "failed to create person")
			return
		}

		c.JSON(http.StatusOK, Response{
			Status: http.StatusText(http.StatusOK),
			Person: createdPerson,
		})
	}
}
