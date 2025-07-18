package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/auth"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
)

func (r Routes) authroutes() {
	// register required services
	command := auth.NewAuthCommandService(
		repositories.NewUserRepository(r.services.DB),
		repositories.NewOtpRepository(r.services.DB),
		r.services.Token,
		r.services.Logger,
		r.services.Mail,
	)

	// register controller
	ctrl := handlers.NewAuthController(command)

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.POST("/login", ctrl.Login)
		auth.POST("/register", ctrl.Register)
	}
}
