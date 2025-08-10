package handlers

import "github.com/gin-gonic/gin"

type VendorController struct {
}

func NewVendorController() *VendorController {
	return &VendorController{}
}

func (ctrl VendorController) Index(c *gin.Context) {

}

func (ctrl VendorController) Store(c *gin.Context) {

}

func (ctrl VendorController) UpdateLogo(c *gin.Context) {

}

func (ctrl VendorController) UpdateBanner(c *gin.Context) {

}

func (ctrl VendorController) Delete(c *gin.Context) {

}
