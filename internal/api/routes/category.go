package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/categories"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
)

func (r Routes) categoryroutes() {
	// register required services
	query := categories.NewCategoryQuery(repositories.NewCategoryRepository(r.services.DB))

	// register controller
	ctrl := handlers.NewCategoryController(query)

	// register routes
	category := r.engine.Group("categories")
	{
		category.GET("/", ctrl.Index)
		category.POST("/", ctrl.Store)
	}
}
