package handlers

import (
	"net/http"

	"github.com/boldd/internal/application/products"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
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

//	@Summary		"generate variant attribute combinations"
//	@Description	"generate variant attribute combinations for a product"
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Schemes
//	@Param		payload	body		products.GenerateCombinationRequest		true	"payload"
//	@Success	200		{array}		[]products.VariantCombinationResponse	"combinations"
//	@Failure	404		{object}	dtos.ErrorResponse						"body"
//	@Failure	500		{object}	dtos.ErrorResponse						"body"
//	@Router		/products/generate-variant-combinations [post]
func (ctrl ProductController) GenerateCombination(c *gin.Context) {
	var payload products.GenerateCombinationRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.command.GenerateCombinations(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, response)
}
