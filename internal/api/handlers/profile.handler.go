package handlers

import (
	"net/http"

	"github.com/boldd/internal/application/profile"
	"github.com/boldd/internal/domain/common"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	command profile.IProfileCommand
}

func NewProfileController(command profile.IProfileCommand) *ProfileController {
	return &ProfileController{command}
}

//	@Summary		"get user profile"
//	@Description	"get user profile information"
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Schemes
//	@Failure	403	{object}	dtos.ErrorResponse	"body"
//	@Failure	500	{object}	dtos.ErrorResponse	"body"
//	@Security	BearerAuth
//	@Router		/profile [get]
func (ctrl ProfileController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, common.GetAuthUser(c))
}

//	@Summary		"change user password"
//	@Description	"change authenticated user password"
//	@Tags			Profile
//	@Accept			json
//	@Produce		json
//	@Schemes
//	@Param		payload	body		profile.ChangePasswordRequest	true	"Change password payload"
//	@Failure	403		{object}	dtos.ErrorResponse				"body"
//	@Failure	400		{object}	dtos.ErrorResponse				"body"
//	@Failure	500		{object}	dtos.ErrorResponse				"body"
//	@Security	BearerAuth
//	@Router		/profile/change-password [post]
func (ctrl ProfileController) ChangePassword(c *gin.Context) {
	user := common.GetAuthUser(c)
	var payload profile.ChangePasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.command.ChangePassword(user, &payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, nil)
}
