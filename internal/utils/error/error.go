package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"status": http.StatusText(statusCode),
		"error":  message,
	})
}
