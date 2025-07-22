package handlers

import (
	"fmt"
	"net/http"

	"github.com/boldd/internal/domain/common"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

// @Summary		"get user profile"
// @Description	"get user profile information"
// @Tags			Profile
// @Accept			json
// @Produce		json
// @Schemes
// @Failure	403			{object}	dtos.ErrorResponse	"body"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Security BearerAuth
// @Router		/profile/ [get]
func (ctrl ProfileController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, common.GetAuthUser(c))
}

// @Summary		"change user password"
// @Description	"change authenticated user password"
// @Tags			Profile
// @Accept			json
// @Produce		json
// @Schemes
// @Failure	403			{object}	dtos.ErrorResponse	"body"
// @Failure	500			{object}	dtos.ErrorResponse	"body"
// @Security BearerAuth
// @Router		/profile/change-password [post]
func (ctrl ProfileController) ChangePassword(c *gin.Context) {
	user := common.GetAuthUser(c)
	fmt.Println(user)
}
