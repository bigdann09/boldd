package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/go-swagno/swagno/components/endpoint"
)

func (r Routes) authroutes() {
	// TODO: register all required services (caching, commands etc..)

	// register controller
	ctrl := handlers.NewAuthController()

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.GET("/register", ctrl.Register)
	}

	// register endpoints
	r.swagger.AddEndpoints([]*endpoint.EndPoint{
		endpoint.New(
			endpoint.GET,
			"/register",
			endpoint.WithTags("auth"),
			endpoint.WithDescription("Register a new user"),
		),
	})
}
