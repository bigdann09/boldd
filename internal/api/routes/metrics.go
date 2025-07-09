package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (r *Routes) metricsroute() {
	r.engine.GET("/metrics", gin.WrapH(promhttp.Handler())) // register metrics
}
