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

// Request структура запроса на создание человека
// @Description Структура запроса для создания нового человека
type Request struct {
	// Name имя человека
	// @Description Имя человека
	// @Required
	Name string `json:"name" binding:"required"`

	// Surname фамилия человека
	// @Description Фамилия человека
	// @Required
	Surname string `json:"surname" binding:"required"`

	// Patronymic отчество человека
	// @Description Отчество человека (опционально)
	Patronymic *string `json:"patronymic"`
}

// Response структура ответа на создание человека
// @Description Структура ответа при создании нового человека
type Response struct {
	// Status статус операции
	// @Description HTTP статус операции
	Status string `json:"status"`

	// Error сообщение об ошибке
	// @Description Сообщение об ошибке (если есть)
	Error string `json:"error,omitempty"`

	// Person данные созданного человека
	// @Description Данные созданного человека
	Person *models.Person `json:"person,omitempty"`
}

// PersonCreator интерфейс для создания человека
type PersonCreator interface {
	CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
}

// @Summary Создать нового человека
// @Description Создает нового человека с указанными данными и обогащает их дополнительной информацией
// @Tags person
// @Accept json
// @Produce json
// @Param request body Request true "Данные для создания человека"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /create [post]
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
