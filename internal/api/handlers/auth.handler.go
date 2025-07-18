package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/boldd/internal/application/auth"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthController struct {
	authsrv auth.IAuthCommandService
}

func NewAuthController(authsrv auth.IAuthCommandService) *AuthController {
	return &AuthController{authsrv}
}

// @Summary		"register user"
// @Description	"Registers a new user"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.RegisterRequest	true	"User registration details"
// @Success	200		{object}	auth.AuthResponse		"body"
// @Router		/auth/register [post]
func (ctrl AuthController) Register(c *gin.Context) {
	var payload auth.RegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.authsrv.Register(&payload)
	if err != nil {
		body := err.(map[string]interface{})
		c.JSON(body["code"].(int), gin.H{
			"message": body["error"],
		})
		return
	}

	accessExpiry, _ := strconv.Atoi(viper.GetStringMapString("jwt")["access_expiry"])
	maxAge := time.Hour * 24 * time.Duration(accessExpiry)
	c.SetCookie("Authorization", response.AccessToken, int(maxAge), "", "", false, false)

	c.JSON(http.StatusOK, response)
}

// @Summary		"authorize a user"
// @Description	"Login user"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.LoginRequest	true	"User Login details"
// @Success	200		{object}	auth.AuthResponse		"body"
// @Router		/auth/login [post]
func (ctrl AuthController) Login(c *gin.Context) {
	var payload auth.LoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.authsrv.Login(&payload)
	if err != nil {
		body := err.(map[string]interface{})
		c.JSON(body["code"].(int), gin.H{
			"message": body["error"],
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (ctrl AuthController) Logout(c *gin.Context)                  {}
func (ctrl AuthController) ForgotPassword(c *gin.Context)          {}
func (ctrl AuthController) ResetPassword(c *gin.Context)           {}
func (ctrl AuthController) ResendConfirmationEmail(c *gin.Context) {}
func (ctrl AuthController) VerifyEmail(c *gin.Context)             {}
func (ctrl AuthController) GoogleLogin(c *gin.Context)             {}
