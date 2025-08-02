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

func (ctrl ProductController) Store(c *gin.Context) {

}
