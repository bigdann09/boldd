package routes

import "github.com/boldd/internal/api/handlers"

func (r Routes) authroutes() {
	// TODO: register all required services (caching, commands etc..)

	// register controller
	ctrl := handlers.NewAuthController()

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.GET("/register", ctrl.Register)
	}
}
