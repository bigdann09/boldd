package handlers

import (
	"net/http"
	"strings"

	"github.com/boldd/internal/application/subcategories"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type SubCategoryController struct {
	query   subcategories.ISubCategoryQuery
	command subcategories.ISubCategoryCommand
}

func NewSubCategoryController(query subcategories.ISubCategoryQuery, command subcategories.ISubCategoryCommand) *SubCategoryController {
	return &SubCategoryController{query, command}
}

// @Summary		"get all subcategories"
// @Description	"get all product subcategories"
// @Tags			Subcategories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		page		query		int					false	"page number"
// @Param		page_size	query		int					false	"page data size"
// @Param		sort_by		query		string				false	"sort by"
// @Param		order		query		string				false	"order"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Router		/subcategories [get]
func (ctrl SubCategoryController) Index(c *gin.Context) {
	var filter dtos.SubCategoryQueryFilter
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

// @Summary		"get a single subcategory"
// @Description	"retrieve a single subcategory from database"
// @Tags			Subcategories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id	path		string						true	"subcategory id"
// @Success	200	{object}	dtos.SubCategoryResponse	"category"
// @Failure	404	{object}	dtos.ErrorResponse			"body"
// @Failure	500	{object}	dtos.ErrorResponse			"body"
// @Router		/subcategories/{id} [get]
func (ctrl SubCategoryController) Show(c *gin.Context) {
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

// @Summary		"product subcategories"
// @Description	"product subcategories"
// @Tags			Subcategories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		subcategories.CreateSubCategoryRequest	true	"Create subcategory payload"
// @Success	201		{null}		null									"no response"
// @Failure	500		{object}	dtos.ErrorResponse						"body"
// @Router		/subcategories [post]
func (ctrl SubCategoryController) Store(c *gin.Context) {
	var payload subcategories.CreateSubCategoryRequest
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

	c.Status(http.StatusCreated)
}

// @Summary		"update a single subcategory"
// @Description	"update a single subcategory in the database"
// @Tags			Subcategories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id		path		string									true	"category id"
// @Param		payload	body		subcategories.UpdateSubCategoryRequest	true	"category id"
// @Failure	404		{object}	dtos.ErrorResponse						"body"
// @Failure	500		{object}	dtos.ErrorResponse						"body"
// @Router		/subcategories/{id} [put]
func (ctrl SubCategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	if strings.EqualFold(id, "") {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "ID invalid or not present", Status: http.StatusBadRequest})
	}

	var payload subcategories.UpdateSubCategoryRequest
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

	c.Status(http.StatusOK)
}

// @Summary		"get a single subcategory"
// @Description	"retrieve a single subcategory from database"
// @Tags			Subcategories
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id	path		string				true	"subcategory id"
// @Failure	404	{object}	dtos.ErrorResponse	"body"
// @Failure	500	{object}	dtos.ErrorResponse	"body"
// @Router		/subcategories/{id} [delete]
func (ctrl SubCategoryController) Delete(c *gin.Context) {
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

	c.Status(http.StatusOK)
}
