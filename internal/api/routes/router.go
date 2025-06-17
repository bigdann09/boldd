package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	engine *gin.Engine
	// logger
	// cache
	// services
}

func NewRouter(engine *gin.Engine) *Routes {
	return &Routes{engine}
}

func (router *Routes) SetupRoutes() *gin.Engine {
	router.authroutes()
	router.healthroute()
	return router.engine
}
