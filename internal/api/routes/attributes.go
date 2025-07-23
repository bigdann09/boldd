package routes

import (
	"github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/application/attributes"
)

func (r Routes) attributesroutes() {
	// register required command services
	command := attributes.NewAttributeCommand()

	ctrl := handlers.NewAttributeController(command)

	attributes := r.engine.Group("attributes/")
	{
		attributes.GET("", ctrl.Index)
		attributes.POST("", ctrl.Store)
	}
}
