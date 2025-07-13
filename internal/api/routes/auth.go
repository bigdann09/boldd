package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/auth"
	"github.com/boldd/internal/infrastructure/persistence/repository"
)

func (r Routes) authroutes() {
	// register required services
	userRepository := repository.NewUserRepository(r.services.DB)
	command := auth.NewAuthCommandService(userRepository)

	// register controller
	ctrl := handlers.NewAuthController(command)

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.POST("/login", ctrl.Login)
		auth.POST("/register", ctrl.Register)
	}
}
