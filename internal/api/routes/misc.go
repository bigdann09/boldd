package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-swagno/swagno-gin/swagger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (r *Routes) miscelleneousroutes() {
	r.engine.GET("/metrics", gin.WrapH(promhttp.Handler())) // register metrics
	r.engine.GET("/swagger/*any", swagger.SwaggerHandler(r.swagger.MustToJson()))
}
