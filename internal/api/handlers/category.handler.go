package handlers

import (
	"net/http"
	"strings"

	"github.com/boldd/internal/application/categories"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	query   categories.ICategoryQuery
	command categories.ICategoryCommand
}

func NewCategoryController(query categories.ICategoryQuery, command categories.ICategoryCommand) *CategoryController {
	return &CategoryController{query, command}
}

// @Summary		"get all categories"
// @Description	"get all product categories"
// @Tags			Categories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		page		query		int					false	"page number"
// @Param		page_size	query		int					false	"page data size"
// @Param		sort_by		query		string				false	"sort by"
// @Param		order		query		string				false	"order"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Router		/categories [get]
func (ctrl CategoryController) Index(c *gin.Context) {
	var filter dtos.CategoryQueryFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.query.GetAll(&filter)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		"get a single category"
// @Description	"retrieve a single category from database"
// @Tags			Categories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id		path		string					true	"category id"
// @Success 200 	{object}	dtos.CategoryResponse									"category"
// @Failure	404		{object}	dtos.ErrorResponse					"body"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/categories/{id} [get]
func (ctrl CategoryController) Show(c *gin.Context) {
	id := c.Param("id")
	if strings.EqualFold(id, "") {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "ID invalid or not present", Status: http.StatusBadRequest})
	}

	response, err := ctrl.query.Get(id)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		"product categories"
// @Description	"product categories"
// @Tags			Categories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		categories.CreateCategoryRequest	true	"Create category payload"
// @Success 201		""			""									"no response"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/categories [post]
func (ctrl CategoryController) Store(c *gin.Context) {
	var payload categories.CreateCategoryRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.command.Create(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// @Summary		"update a single category"
// @Description	"update a single category in the database"
// @Tags			Categories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id		path		string					true	"category id"
// @Param		payload		body		categories.UpdateCategoryRequest		true	"category id"
// @Failure	404		{object}	dtos.ErrorResponse					"body"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/categories/{id} [put]
func (ctrl CategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	if strings.EqualFold(id, "") {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "ID invalid or not present", Status: http.StatusBadRequest})
	}

	var payload categories.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.command.Update(id, &payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary		"get a single category"
// @Description	"retrieve a single category from database"
// @Tags			Categories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id		path		string					true	"category id"
// @Failure	404		{object}	dtos.ErrorResponse					"body"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/categories/{id} [delete]
func (ctrl CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	if strings.EqualFold(id, "") {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "ID invalid or not present", Status: http.StatusBadRequest})
	}

	err := ctrl.command.Delete(id)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, nil)
}
