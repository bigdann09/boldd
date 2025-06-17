package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Routes) healthroute() {
	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "server healthy",
		})
	})
}
