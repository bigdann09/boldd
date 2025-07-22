package common

import (
	"github.com/boldd/internal/domain/dtos"
	"github.com/gin-gonic/gin"
)

func GetAuthUser(c *gin.Context) *dtos.UserResponse {
	user := c.MustGet("user").(dtos.UserResponse)
	return &user
}
