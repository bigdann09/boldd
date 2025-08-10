package routes

import "github.com/boldd/internal/api/handlers"

func (r Routes) vendorroutes() {
	ctrl := handlers.NewVendorController()

	vendors := r.engine.Group("vendors/")
	vendors.Use(r.middlewares.Auth())
	{
		vendors.GET("", r.middlewares.Admin(), ctrl.Index)
		vendors.POST("", ctrl.Store)
		vendors.DELETE("/:id", r.middlewares.Admin(), ctrl.Delete)
		vendors.PUT("/:id/upload/logo", r.middlewares.Vendor(), ctrl.UpdateLogo)
		vendors.PUT("/:id/upload/banner", r.middlewares.Vendor(), ctrl.UpdateBanner)
	}
}
