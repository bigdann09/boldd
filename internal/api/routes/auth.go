package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/auth"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
)

func (r Routes) authroutes() {
	// register required services
	userRepository := repositories.NewUserRepository(r.services.DB)
	command := auth.NewAuthCommandService(userRepository, r.services.Token, r.services.Logger)

	// register controller
	ctrl := handlers.NewAuthController(command)

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.POST("/login", ctrl.Login)
		auth.POST("/register", ctrl.Register)
	}
}
