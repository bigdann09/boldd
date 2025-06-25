package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (r *Routes) miscelleneousroute() {
	r.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
