package routes

import (
	"github.com/boldd/internal/api/handlers"
)

func (r Routes) profileroutes() {
	// register required services

	// register controller
	ctrl := handlers.NewProfileController()

	// register routes
	profile := r.engine.Group("profile")
	profile.Use(r.middlewares.Auth())
	{
		profile.GET("/", ctrl.Show)
	}
}
