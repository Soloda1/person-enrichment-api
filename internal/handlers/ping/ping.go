package ping

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Проверка работоспособности API
// @Description Возвращает статус работоспособности API и текущее время
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{} "Успешный ответ с сообщением pong и временем"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"timestamp": time.Now(),
	})
}
