package middlewares

import (
	"strconv"
	"time"

	"github.com/boldd/internal/infrastructure/monitoring"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) Metrics(metrics *monitoring.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		metrics.RequestsTotal.WithLabelValues(
			c.Request.Method,
			c.Request.URL.Path,
			strconv.Itoa(c.Writer.Status()),
		).Inc()

		metrics.RequestLatency.WithLabelValues(
			c.Request.Method,
			c.Request.URL.Path,
		).Observe(time.Since(start).Seconds())
	}
}
