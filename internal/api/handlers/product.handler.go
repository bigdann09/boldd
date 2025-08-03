package handlers

import (
	"github.com/boldd/internal/application/products"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	command products.IProductCommand
}

func NewProductController(command products.IProductCommand) *ProductController {
	return &ProductController{command}
}

func (ctrl ProductController) Index(c *gin.Context) {

}

//	@Summary		"store a product"
//	@Description	"store a new product"
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Schemes
//	@Param		payload	body		products.CreateProductRequest	true	"product payload"
//	@Success	201		{string}	string							"No Content"
//	@Failure	404		{object}	dtos.ErrorResponse				"body"
//	@Failure	500		{object}	dtos.ErrorResponse				"body"
//	@Router		/products [post]
func (ctrl ProductController) Store(c *gin.Context) {

}
