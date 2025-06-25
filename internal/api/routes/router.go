package routes

import (
	"github.com/boldd/internal/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-swagno/swagno"
)

type Routes struct {
	engine   *gin.Engine
	services *services.Service
	swagger  *swagno.Swagger
}

func NewRouter(engine *gin.Engine, services *services.Service, swagger *swagno.Swagger) *Routes {
	return &Routes{engine, services, swagger}
}

func (router *Routes) SetupRoutes() *gin.Engine {
	router.authroutes()
	router.healthroute()
	router.miscelleneousroutes()
	return router.engine
}
