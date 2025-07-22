package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.services.Logger.Info("Retrieve bearer token from header")
		value := c.GetHeader("Authorization")

		if strings.EqualFold(value, "") {
			m.services.Logger.Error("jwt token not provided in header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized access to resource",
			})
		}

		m.services.Logger.Info("parse token")
		var token string
		token = value
		if strings.Contains(value, "Bearer") || strings.Contains(value, "bearer") {
			token = strings.Split(value, " ")[1]
		}
		if strings.EqualFold(token, "") {
			m.services.Logger.Error("invalid jwt token provided in header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized access to resource",
			})
		}

		m.services.Logger.Info("validate jwt token from header")
		claims, err := m.services.Token.ValidateToken(token)
		if err != nil {
			m.services.Logger.Error("jwt token not valid or expired", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized access to resource",
			})
		}

		m.services.Logger.Info("retrieve and cache user information")
		repository := repositories.NewUserRepository(m.services.DB)
		cache := cache.NewCache[dtos.UserResponse](m.services.Redis, c.Request.Context(), time.Duration(time.Second*40))
		user, _ := cache.GetOrSet(fmt.Sprintf("auth_user:%d", claims.Id), func() (dtos.UserResponse, error) {
			return repository.FindByEmail(claims.Email)
		})

		m.services.Logger.Info("store user in context")
		c.Set("user", user)
		c.Set("roles", user.Roles)

		c.Next()
	}
}
