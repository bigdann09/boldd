package routes

import (
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func (r *Routes) swaggerroutes() {
	r.engine.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))
}
