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

// Request структура запроса на обновление человека
// @Description Структура запроса для обновления информации о человеке
type Request struct {
	// ID идентификатор человека
	// @Description ID человека для обновления
	// @Required
	ID int `json:"id" binding:"required"`

	// Name новое имя человека
	// @Description Новое имя человека (опционально)
	Name *string `json:"name,omitempty"`

	// Surname новая фамилия человека
	// @Description Новая фамилия человека (опционально)
	Surname *string `json:"surname,omitempty"`

	// Patronymic новое отчество человека
	// @Description Новое отчество человека (опционально)
	Patronymic *string `json:"patronymic,omitempty"`

	// Age новый возраст человека
	// @Description Новый возраст человека (опционально)
	Age *int `json:"age,omitempty"`

	// Gender новый пол человека
	// @Description Новый пол человека (опционально)
	Gender *string `json:"gender,omitempty"`

	// National новые национальности
	// @Description Новый список национальностей (опционально)
	National *[]string `json:"national,omitempty"`
}

// Response структура ответа на обновление человека
// @Description Структура ответа при обновлении информации о человеке
type Response struct {
	// Status статус операции
	// @Description HTTP статус операции
	Status string `json:"status"`

	// Error сообщение об ошибке
	// @Description Сообщение об ошибке (если есть)
	Error string `json:"error,omitempty"`

	// Person обновленные данные человека
	// @Description Обновленные данные человека
	Person *models.Person `json:"person,omitempty"`
}

// PersonUpdater интерфейс для обновления данных человека
type PersonUpdater interface {
	UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
}

// @Summary Обновить информацию о человеке
// @Description Обновляет информацию о человеке по его ID
// @Tags person
// @Accept json
// @Produce json
// @Param request body Request true "Данные для обновления"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /update [put]
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
