package routes

import (
	"context"
	"time"

	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/vendors"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
)

func (r Routes) vendorroutes() {
	command := vendors.NewVendorCommand(
		r.services.Logger,
		repositories.NewVendorRepository(r.services.DB),
		cache.NewCache[dtos.VendorResponse](
			r.services.Redis,
			context.Background(),
			time.Duration(time.Second*10),
		),
	)

	ctrl := handlers.NewVendorController(command)

	vendors := r.engine.Group("vendors/")
	vendors.Use(r.middlewares.Auth())
	{
		vendors.POST("", ctrl.Store)
		vendors.PUT(
			"/:id/upload/logo",
			r.middlewares.Vendor(),
			ctrl.UpdateLogo,
		)
		vendors.PUT(
			"/:id/upload/banner",
			r.middlewares.Vendor(),
			ctrl.UpdateBanner,
		)
		vendors.GET(
			"",
			r.middlewares.Admin(),
			ctrl.Index,
		)
		vendors.DELETE(
			"/:id",
			r.middlewares.Admin(),
			ctrl.Delete,
		)

	}
}
