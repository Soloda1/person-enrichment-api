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

// Response структура ответа со списком людей
// @Description Структура ответа со списком людей и дополнительной информацией
type Response struct {
	// Status статус операции
	// @Description HTTP статус операции
	Status string `json:"status"`

	// Error сообщение об ошибке
	// @Description Сообщение об ошибке (если есть)
	Error string `json:"error,omitempty"`

	// Persons список людей
	// @Description Список найденных людей
	Persons []*models.Person `json:"person,omitempty"`
}

// PersonsProvider интерфейс для получения списка людей
type PersonsProvider interface {
	GetAllPersons(ctx context.Context, filter models.PersonFilter) ([]*models.Person, error)
}

// @Summary Получить список людей
// @Description Получает список людей с возможностью фильтрации и пагинации
// @Tags person
// @Accept json
// @Produce json
// @Param name query string false "Фильтр по имени"
// @Param surname query string false "Фильтр по фамилии"
// @Param patronymic query string false "Фильтр по отчеству"
// @Param gender query string false "Фильтр по полу"
// @Param national query string false "Фильтр по национальности"
// @Param min_age query int false "Минимальный возраст"
// @Param max_age query int false "Максимальный возраст"
// @Param limit query int false "Количество записей на странице"
// @Param page query int false "Номер страницы"
// @Success 200 {object} Response
// @Failure 500 {object} utils.ErrorResponse
// @Router /persons [get]
func New(log *logger.Logger, service PersonsProvider, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("GetAllPersons called")

		filter := models.PersonFilter{}

		if name := c.Query("name"); name != "" {
			filter.Name = &name
		}
		if surname := c.Query("surname"); surname != "" {
			filter.Surname = &surname
		}
		if patronymic := c.Query("patronymic"); patronymic != "" {
			filter.Patronymic = &patronymic
		}
		if gender := c.Query("gender"); gender != "" {
			filter.Gender = &gender
		}
		if national := c.Query("national"); national != "" {
			filter.National = &national
		}
		if minAgeStr := c.Query("min_age"); minAgeStr != "" {
			if minAge, err := strconv.Atoi(minAgeStr); err == nil {
				filter.MinAge = &minAge
			}
		}
		if maxAgeStr := c.Query("max_age"); maxAgeStr != "" {
			if maxAge, err := strconv.Atoi(maxAgeStr); err == nil {
				filter.MaxAge = &maxAge
			}
		}

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
