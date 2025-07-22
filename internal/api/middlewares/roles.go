package middlewares

import (
	"net/http"
	"slices"

	"github.com/boldd/internal/domain/dtos"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.services.Logger.Info("retrieve authenticated user")
		roles := c.MustGet("roles").([]string)
		if !slices.Contains(roles, "admin") {
			m.services.Logger.Error("forbidden access to admin resource")
			c.AbortWithStatusJSON(http.StatusForbidden, dtos.ErrorResponse{
				Status:  http.StatusForbidden,
				Message: "forbidden access to resource",
			})
		}
	}
}

func (m *Middleware) Merchant() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.services.Logger.Info("retrieve authenticated user")
		roles := c.MustGet("roles").([]string)
		if !slices.Contains(roles, "merchant") {
			m.services.Logger.Error("forbidden access to merchant resource")
			c.AbortWithStatusJSON(http.StatusForbidden, dtos.ErrorResponse{
				Status:  http.StatusForbidden,
				Message: "forbidden access to resource",
			})
		}
	}
}
