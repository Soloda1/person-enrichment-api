package requestLogger

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Logger Middleware enabled")

		entry := log.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("remote_address", c.ClientIP()))

		start := time.Now()

		c.Next()

		latency := time.Since(start)

		entry.Info("request completed",
			slog.String("latency", latency.String()),
			slog.Int("status", c.Writer.Status()))
	}
}
