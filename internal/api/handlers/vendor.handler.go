package handlers

import (
	"net/http"

	"github.com/boldd/internal/application/vendors"
	"github.com/boldd/internal/domain/common"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type VendorController struct {
	command vendors.IVendorCommand
}

func NewVendorController(command vendors.IVendorCommand) *VendorController {
	return &VendorController{command: command}
}

// @Summary		"get all vendors"
// @Description	"get all vendors"
// @Tags			Vendors
// @Accept			json
// @Produce		json
// @Schemes
// @Param		page		query		int					false	"page number"
// @Param		page_size	query		int					false	"page data size"
// @Param		sort_by		query		string				false	"sort by"
// @Param		order		query		string				false	"order"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Security	BearerAuth
// @Router		/vendors [get]
func (ctrl VendorController) Index(c *gin.Context) {
	var filter dtos.VendorQueryFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		validator.GetErrors(c, err)
		return
	}
}

// @Summary		"store vendors"
// @Description	"store vendors"
// @Tags			Vendors
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		vendors.CreateVendorRequest	true	"Create vendor payload"
// @Success	201		{string}	null								"No Content"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Security	BearerAuth
// @Router		/vendors [post]
func (ctrl VendorController) Store(c *gin.Context) {
	user := common.GetAuthUser(c)
	var payload vendors.CreateVendorRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.command.Create(user, &payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary		"Update vendor logo"
// @Description	"Update vendor logo"
// @Tags			Vendors
// @Accept			mpfd
// @Produce		json
// @Schemes
// @Param		image	formData		file	true	"Logo image file"
// @Success	201		{string}	null								"No Content"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Security	BearerAuth
// @Router		/vendors/{id}/upload/logo [put]
func (ctrl VendorController) UpdateLogo(c *gin.Context) {

}

func (ctrl VendorController) UpdateBanner(c *gin.Context) {

}

func (ctrl VendorController) Delete(c *gin.Context) {

}
