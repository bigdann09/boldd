package handlers

import (
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type VendorController struct {
}

func NewVendorController() *VendorController {
	return &VendorController{}
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
// @Router		/vendors [get]
func (ctrl VendorController) Index(c *gin.Context) {
	var filter dtos.VendorQueryFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		validator.GetErrors(c, err)
		return
	}
}

func (ctrl VendorController) Store(c *gin.Context) {

}

func (ctrl VendorController) UpdateLogo(c *gin.Context) {

}

func (ctrl VendorController) UpdateBanner(c *gin.Context) {

}

func (ctrl VendorController) Delete(c *gin.Context) {

}
