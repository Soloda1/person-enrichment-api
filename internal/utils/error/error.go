package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse структура ответа с ошибкой
// @Description Стандартная структура ответа при возникновении ошибки
type ErrorResponse struct {
	// Status HTTP статус
	// @Description HTTP статус ошибки
	Status string `json:"status"`

	// Error сообщение об ошибке
	// @Description Описание ошибки
	Error string `json:"error"`
}

func SendError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Status: http.StatusText(statusCode),
		Error:  message,
	})
}
