package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (m *Middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Logging informasi lengkap
		m.Log.Info("Request processed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
		)
	}
}
