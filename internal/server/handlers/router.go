package handlers

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/server/handlers/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) *gin.Engine {
	fmt.Println("Registering application routes...")

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	engine.MaxMultipartMemory = 8 << 20 // 8MB

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Boldd API!")
	})

	// Register application routes
	auth.Routes(engine)

	return engine
}
