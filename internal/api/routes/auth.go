package routes

import "github.com/boldd/internal/api/handlers"

func (r Routes) authroutes() {
	// register all services required
	// usersrv := NewAuthService(database, logger, caching)

	// register controller
	ctrl := handlers.NewAuthController()

	// register routes
	auth := r.engine.Group("auth")
	{
		auth.GET("/register", ctrl.Register)
	}

}
