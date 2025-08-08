package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/products"
)

func (r *Routes) productroutes() {
	// register command service
	command := products.NewProductCommand()

	// register controller
	ctrl := handlers.NewProductController(command)

	// register product routes
	product := r.engine.Group("products/")
	{
		product.GET("", ctrl.Index)
		product.POST("", ctrl.Store)
		product.POST("/generate-variant-combinations", ctrl.GenerateCombination)
	}
}
