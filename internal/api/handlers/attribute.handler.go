package handlers

import (
	"net/http"
	"strings"

	"github.com/boldd/internal/application/attributes"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type AttributeController struct {
	query   attributes.IAttributeQuery
	command attributes.IAttributeCommand
}

func NewAttributeController(query attributes.IAttributeQuery, command attributes.IAttributeCommand) *AttributeController {
	return &AttributeController{query, command}
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
	var filter dtos.AttributeQueryFilter
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

// @Summary		"get a single attribute"
// @Description	"retrieve a single attribute from database"
// @Tags			Attributes
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id	path		string					true	"attribute id"
// @Success	200	{object}	dtos.AttributeResponse	"attribute"
// @Failure	404	{object}	dtos.ErrorResponse		"body"
// @Failure	500	{object}	dtos.ErrorResponse		"body"
// @Router		/attributes/{id} [get]
func (ctrl AttributeController) Show(c *gin.Context) {
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

// @Summary		"get a single attribute"
// @Description	"retrieve a single attribute from database"
// @Tags			Attributes
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		attributes.CreateAttributeRequest	true	"Create attribute payload"
// @Success	201		""			""									"no response"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/attributes [post]
func (ctrl AttributeController) Store(c *gin.Context) {
	var payload attributes.CreateAttributeRequest
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

// @Summary		"update a single attribute"
// @Description	"update a single attribute in the database"
// @Tags			Attributes
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id		path		string								true	"attribute id"
// @Param		payload	body		attributes.UpdateAttributeRequest	true	"attribute id"
// @Failure	404		{object}	dtos.ErrorResponse					"body"
// @Failure	500		{object}	dtos.ErrorResponse					"body"
// @Router		/attributes/{id} [put]
func (ctrl AttributeController) Update(c *gin.Context) {
	id := c.Param("id")
	if strings.EqualFold(id, "") {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "ID invalid or not present", Status: http.StatusBadRequest})
	}

	var payload attributes.UpdateAttributeRequest
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

// @Summary		"get a single attribute"
// @Description	"retrieve a single attribute from database"
// @Tags			Attributes
// @Accept			json
// @Produce		json
// @Schemes
// @Param		id	path		string				true	"attribute id"
// @Failure	404	{object}	dtos.ErrorResponse	"body"
// @Failure	500	{object}	dtos.ErrorResponse	"body"
// @Router		/attributes/{id} [delete]
func (ctrl AttributeController) Delete(c *gin.Context) {
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
