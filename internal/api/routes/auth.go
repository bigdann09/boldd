package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/auth"
)

func (r Routes) authroutes() {
	// register required services
	command := auth.NewAuthCommandService()

	// register controller
	ctrl := handlers.NewAuthController(command)

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.POST("/login", ctrl.Login)
		auth.POST("/register", ctrl.Register)
	}
}
