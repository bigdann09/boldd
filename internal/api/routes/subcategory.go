package routes

import (
	"context"
	"time"

	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/subcategories"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
)

func (r Routes) subcategoryroutes() {
	// register required services
	query := subcategories.NewSubCategoryQuery(
		r.services.Logger,
		repositories.NewSubCategoryRepository(r.services.DB),
		cache.NewCache[*dtos.SubCategoryResponse](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
		cache.NewCache[utils.PaginationResponse[dtos.SubCategoryResponse]](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
	)

	command := subcategories.NewSubCategoryCommand(
		r.services.Logger,
		repositories.NewCategoryRepository(r.services.DB),
		repositories.NewSubCategoryRepository(r.services.DB),
		cache.NewCache[entities.SubCategory](
			r.services.Redis,
			context.Background(),
			time.Minute*5,
		),
	)

	// register controller
	ctrl := handlers.NewSubCategoryController(query, command)

	// register routes
	subcategory := r.engine.Group("subcategories")
	{
		subcategory.GET("/", ctrl.Index)
		subcategory.POST("/", ctrl.Store)
		subcategory.GET("/:id", ctrl.Show)
		subcategory.PUT("/:id", ctrl.Update)
		subcategory.DELETE("/:id", ctrl.Delete)
	}
}
