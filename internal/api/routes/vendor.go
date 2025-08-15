package routes

import "github.com/boldd/internal/api/handlers"

func (r Routes) vendorroutes() {
	ctrl := handlers.NewVendorController()

	vendors := r.engine.Group("vendors/")
	vendors.Use(r.middlewares.Auth())
	{
		vendors.POST("", ctrl.Store)
		vendors.Use(r.middlewares.Vendor())
		{
			vendors.PUT("/:id/upload/logo", ctrl.UpdateLogo)
			vendors.PUT("/:id/upload/banner", ctrl.UpdateBanner)
		}

		vendors.Use(r.middlewares.Admin())
		{
			vendors.GET("", ctrl.Index)
			vendors.DELETE("/:id", ctrl.Delete)
		}

	}
}
