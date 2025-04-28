package updatePerson

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
	ID         int       `json:"id" binding:"required"`
	Name       *string   `json:"name,omitempty"`
	Surname    *string   `json:"surname,omitempty"`
	Patronymic *string   `json:"patronymic,omitempty"`
	Age        *int      `json:"age,omitempty"`
	Gender     *string   `json:"gender,omitempty"`
	National   *[]string `json:"national,omitempty"`
}

type Response struct {
	Status string         `json:"status"`
	Error  string         `json:"error,omitempty"`
	Person *models.Person `json:"person,omitempty"`
}

type PersonUpdater interface {
	UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
}

func New(log *logger.Logger, service PersonUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("UpdatePerson called")

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Debug("failed to bind request body", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusBadRequest, "invalid request body")
			return
		}

		personModel := &models.Person{
			ID: req.ID,
		}

		if req.Name != nil {
			personModel.Name = *req.Name
		}

		if req.Surname != nil {
			personModel.Surname = *req.Surname
		}

		if req.Patronymic != nil {
			personModel.Patronymic = req.Patronymic
		}

		if req.Age != nil {
			personModel.Age = *req.Age
		}

		if req.Gender != nil {
			personModel.Gender = *req.Gender
		}

		if req.National != nil {
			personModel.National = *req.National
		}

		updatedPerson, err := service.UpdatePerson(c.Request.Context(), personModel)
		if err != nil {
			log.Debug("failed to update person", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, "failed to update person")
			return
		}

		c.JSON(http.StatusOK, Response{
			Status: http.StatusText(http.StatusOK),
			Person: updatedPerson,
		})
	}
}
