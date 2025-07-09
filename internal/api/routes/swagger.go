package routes

import (
	"github.com/boldd/internal/api/docs"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func (r *Routes) swaggerroutes() {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Boldd API"
	docs.SwaggerInfo.Description = "Boldd Ecommerce API"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = viper.GetString("application.url")

	r.engine.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))
}
