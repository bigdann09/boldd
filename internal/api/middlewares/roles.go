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

func (m *Middleware) Vendor() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.services.Logger.Info("retrieve authenticated user")
		roles := c.MustGet("roles").([]string)
		if !slices.Contains(roles, "vendor") {
			m.services.Logger.Error("forbidden access to vendor resource")
			c.AbortWithStatusJSON(http.StatusForbidden, dtos.ErrorResponse{
				Status:  http.StatusForbidden,
				Message: "forbidden access to resource",
			})
		}
	}
}

func (m *Middleware) Roles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		m.services.Logger.Info("retrieve authenticated user")
		user_roles := c.MustGet("roles").([]string)
		for _, role := range user_roles {
			if !slices.Contains(roles, role) {
				m.services.Logger.Error("forbidden access to vendor resource")
				c.AbortWithStatusJSON(http.StatusForbidden, dtos.ErrorResponse{
					Status:  http.StatusForbidden,
					Message: "forbidden access to resource",
				})
			}
		}
	}
}
