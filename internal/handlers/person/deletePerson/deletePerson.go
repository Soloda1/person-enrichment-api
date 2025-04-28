package deletePerson

import (
	"context"
	"log/slog"
	"net/http"
	utils "person-enrichment-api/internal/utils/error"
	"person-enrichment-api/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// Request структура запроса на удаление человека
// @Description Структура запроса для удаления человека по ID
type Request struct {
	// PersonId идентификатор человека
	// @Description ID человека для удаления
	PersonId int `uri:"id" binding:"required"`
}

// Response структура ответа на удаление человека
// @Description Структура ответа при удалении человека
type Response struct {
	// Status статус операции
	// @Description HTTP статус операции
	Status string `json:"status"`

	// Error сообщение об ошибке
	// @Description Сообщение об ошибке (если есть)
	Error string `json:"error,omitempty"`
}

// PersonDeleter интерфейс для удаления человека
type PersonDeleter interface {
	DeletePerson(ctx context.Context, personId int) error
}

// @Summary Удалить человека
// @Description Удаляет человека по его ID
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Success 200 {object} Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /delete/{id} [delete]
func New(log *logger.Logger, service PersonDeleter) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("DeletePerson called")

		var req Request
		if err := c.ShouldBindUri(&req); err != nil {
			log.Debug("personId is invalid", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusBadRequest, "invalid person id")
			return
		}

		err := service.DeletePerson(c.Request.Context(), req.PersonId)
		if err != nil {
			log.Debug("failed to delete person", slog.String("error", err.Error()))
			utils.SendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, Response{
			Status: http.StatusText(http.StatusOK),
		})
	}
}
