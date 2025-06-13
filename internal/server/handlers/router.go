package handlers

import (
	"fmt"

	"github.com/boldd/internal/api/handlers/auth"
)

func RegisterRoutes() {
	fmt.Println("Registering routes...")

	// Register the auth routes
	auth.Routes()

	// You can add more route registrations here as needed
	// For example:
	// user.RegisterRoutes()
	// product.RegisterRoutes()
	// order.RegisterRoutes()
}
