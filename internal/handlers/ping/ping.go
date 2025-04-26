package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"timestamp": time.Now(),
	})
}
