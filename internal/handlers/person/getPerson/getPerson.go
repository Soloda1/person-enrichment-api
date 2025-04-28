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

// Request структура запроса на получение человека
// @Description Структура запроса для получения информации о человеке по ID
type Request struct {
	// PersonID идентификатор человека
	// @Description ID человека для поиска
	PersonID int `uri:"id" binding:"required"`
}

// Response структура ответа с данными человека
// @Description Структура ответа с информацией о человеке
type Response struct {
	// Status статус операции
	// @Description HTTP статус операции
	Status string `json:"status"`

	// Error сообщение об ошибке
	// @Description Сообщение об ошибке (если есть)
	Error string `json:"error,omitempty"`

	// Person данные человека
	// @Description Данные найденного человека
	Person *models.Person `json:"person,omitempty"`
}

// PersonByIdProvider интерфейс для получения человека по ID
type PersonByIdProvider interface {
	GetPersonByID(ctx context.Context, personId int) (*models.Person, error)
}

// @Summary Получить человека по ID
// @Description Получает информацию о человеке по его идентификатору
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /person/{id} [get]
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
