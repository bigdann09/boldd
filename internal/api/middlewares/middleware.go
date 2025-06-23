package middlewares

import (
	"github.com/boldd/internal/api/services"
	"github.com/boldd/internal/config"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	services *services.Service
	engine   *gin.Engine
}

func NewMiddleware(engine *gin.Engine, services *services.Service) *Middleware {
	return &Middleware{services, engine}
}

func (m *Middleware) Register(cfg *config.Config) {
	m.engine.Use(gzip.Gzip(gzip.DefaultCompression)) // register gzip
	m.engine.Use(m.Cors(&cfg.CorsConfig))            // register cors
}
