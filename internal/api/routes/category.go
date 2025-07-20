package routes

import (
	"context"
	"time"

	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/categories"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
)

func (r Routes) categoryroutes() {
	// register required services
	query := categories.NewCategoryQuery(
		r.services.Logger,
		repositories.NewCategoryRepository(r.services.DB),
		cache.NewCache[utils.PaginationResponse[dtos.CategoryResponse]](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
	)

	// register controller
	ctrl := handlers.NewCategoryController(query)

	// register routes
	category := r.engine.Group("categories")
	{
		category.GET("/", ctrl.Index)
		category.POST("/", ctrl.Store)
	}
}
