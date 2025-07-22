package routes

import (
	"github.com/boldd/internal/api/middlewares"
	"github.com/boldd/internal/api/services"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	engine      *gin.Engine
	services    *services.Service
	middlewares *middlewares.Middleware
}

func NewRouter(engine *gin.Engine, services *services.Service, middlewares *middlewares.Middleware) *Routes {
	return &Routes{engine, services, middlewares}
}

func (router *Routes) SetupRoutes() *gin.Engine {
	router.authroutes()
	router.healthroute()
	router.metricsroute()
	router.swaggerroutes()
	router.profileroutes()
	router.categoryroutes()
	router.subcategoryroutes()
	return router.engine
}
