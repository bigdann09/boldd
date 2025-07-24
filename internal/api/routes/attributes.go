package routes

import (
	"context"
	"time"

	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/attributes"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
)

func (r Routes) attributesroutes() {
	// register required command services
	command := attributes.NewAttributeCommand(
		r.services.Logger,
		repositories.NewAttributeRepository(r.services.DB),
		cache.NewCache[entities.Attribute](
			r.services.Redis,
			context.Background(),
			time.Duration(time.Second*40),
		),
	)

	// register required query services
	query := attributes.NewAttributeQuery(
		r.services.Logger,
		repositories.NewAttributeRepository(r.services.DB),
		cache.NewCache[*dtos.AttributeResponse](
			r.services.Redis,
			context.Background(),
			time.Duration(time.Second*40),
		),
		cache.NewCache[utils.PaginationResponse[dtos.AttributeResponse]](
			r.services.Redis,
			context.Background(),
			time.Duration(time.Second*40),
		),
	)

	ctrl := handlers.NewAttributeController(query, command)

	attributes := r.engine.Group("attributes/")
	{
		attributes.GET("", ctrl.Index)
		attributes.POST("", ctrl.Store)
	}
}
