package handlers

import (
	"net/http"

	"github.com/boldd/internal/application/categories"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	query *categories.CategoryQuery
}

func NewCategoryController(query *categories.CategoryQuery) *CategoryController {
	return &CategoryController{query}
}

// @Summary		"get all categories"
// @Description	"get all product categories"
// @Tags			Category
// @Accept			json
// @Produce		json
// @Schemes
// @Param		page		query		int					false	"page number"
// @Param		page_size	query		int					false	"page data size"
// @Param		sort_by		query		string				false	"sort by"
// @Param		order		query		string				false	"order"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Router		/categories/ [get]
func (ctrl CategoryController) Index(c *gin.Context) {
	var filter dtos.CategoryQueryFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.query.GetAll(&filter)
	if err != nil {
		body := err.(map[string]interface{})
		code := body["code"].(int)
		c.JSON(code, dtos.ErrorResponse{
			Message: body["error"].(string),
			Status:  code,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl CategoryController) Show(c *gin.Context) {

}

// @Summary		"product categories"
// @Description	"product categories"
// @Tags			Category
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		categories.CreateCategoryRequest	true	"Create category payload"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/categories/ [post]
func (ctrl CategoryController) Store(c *gin.Context) {

}

func (ctrl CategoryController) Update(c *gin.Context) {

}

func (ctrl CategoryController) Delete(c *gin.Context) {

}
