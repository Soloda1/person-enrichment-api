package deletePerson

import (
	"context"
	"log/slog"
	"net/http"
	utils "person-enrichment-api/internal/utils/error"
	"person-enrichment-api/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type Request struct {
	PersonId int `uri:"id" binding:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type PersonDeleter interface {
	DeletePerson(ctx context.Context, personId int) error
}

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
