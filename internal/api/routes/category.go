package routes

import (
	"context"
	"time"

	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/categories"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
)

func (r Routes) categoryroutes() {
	// register required services
	query := categories.NewCategoryQuery(
		r.services.Logger,
		repositories.NewCategoryRepository(r.services.DB),
		cache.NewCache[*dtos.CategoryResponse](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
		cache.NewCache[utils.PaginationResponse[dtos.CategoryResponse]](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
	)

	command := categories.NewCategoryCommand(
		r.services.Logger,
		repositories.NewCategoryRepository(r.services.DB),
		cache.NewCache[entities.Category](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
	)

	// register controller
	ctrl := handlers.NewCategoryController(query, command)

	// register routes
	category := r.engine.Group("categories")
	{
		category.GET("/", ctrl.Index)
		category.POST("/", ctrl.Store)
		category.GET("/:uuid", ctrl.Show)
		category.PUT("/:uuid", ctrl.Update)
		category.DELETE("/:uuid", ctrl.Delete)
	}
}
