package handlers

import (
	"github.com/boldd/internal/application/attributes"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type AttributeController struct {
	command attributes.IAttributeCommand
}

func NewAttributeController(command attributes.IAttributeCommand) *AttributeController {
	return &AttributeController{command}
}

// @Summary		"get all product attributes"
// @Description	"get all product attributes"
// @Tags			Attributes
// @Accept			json
// @Produce		json
// @Schemes
// @Param		page		query		int					false	"page number"
// @Param		page_size	query		int					false	"page data size"
// @Param		sort_by		query		string				false	"sort by"
// @Param		order		query		string				false	"order"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Router		/attributes [get]
func (ctrl AttributeController) Index(c *gin.Context) {

}

// @Summary		"get a single attribute"
// @Description	"retrieve a single attribute from database"
// @Tags			Attributes
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		attributes.CreateAttributeRequest	true	"Create attribute payload"
// @Success 201		""			""									"no response"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/attributes [post]
func (ctrl AttributeController) Store(c *gin.Context) {
	var payload attributes.CreateAttributeRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}
}
