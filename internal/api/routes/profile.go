package routes

import (
	"context"
	"time"

	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/profile"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/cache"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
)

func (r Routes) profileroutes() {
	// register required services
	command := profile.NewProfileCommand(
		r.services.Logger,
		repositories.NewUserRepository(r.services.DB),
		cache.NewCache[*dtos.UserResponse](
			r.services.Redis,
			context.Background(),
			time.Second*10,
		),
	)

	// register controller
	ctrl := handlers.NewProfileController(command)

	// register routes
	profile := r.engine.Group("profile/")
	profile.Use(r.middlewares.Auth())
	{
		profile.GET("", ctrl.Show)
		profile.POST("/change-password", ctrl.ChangePassword)
	}
}
