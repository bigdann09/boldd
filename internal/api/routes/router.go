package routes

import (
	"github.com/boldd/internal/api/services"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	engine   *gin.Engine
	services *services.Service
}

func NewRouter(engine *gin.Engine, services *services.Service) *Routes {
	return &Routes{engine, services}
}

func (router *Routes) SetupRoutes() *gin.Engine {
	router.authroutes()
	router.healthroute()
	router.metricsroute()
	router.swaggerroutes()
	return router.engine
}
