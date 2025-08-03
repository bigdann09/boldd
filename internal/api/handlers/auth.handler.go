package handlers

import (
	"net/http"

	"github.com/boldd/internal/application/auth"
	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type Key string

const providerkey Key = "provider"

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
// @Failure	500		{object}	dtos.ErrorResponse		"body"
// @Router		/auth/register [post]
func (ctrl AuthController) Register(c *gin.Context) {
	var payload auth.RegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.authsrv.Register(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary		"authorize a user"
// @Description	"Login user"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.LoginRequest	true	"User Login details"
// @Success	200		{object}	auth.AuthResponse	"body"
// @Failure	400		{object}	dtos.ErrorResponse	"body"
// @Failure	500		{object}	dtos.ErrorResponse	"body"
// @Router		/auth/login [post]
func (ctrl AuthController) Login(c *gin.Context) {
	var payload auth.LoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.authsrv.Login(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		"refresh token"
// @Description	"Refresh user access token"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.RefreshTokenRequest	true	"User registration details"
// @Success	200		{object}	auth.AuthResponse			"body"
// @Failure	500		{object}	dtos.ErrorResponse			"body"
// @Router		/auth/refresh-token [post]
func (ctrl AuthController) RefreshToken(c *gin.Context) {
	var payload auth.RefreshTokenRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	response, err := ctrl.authsrv.RefreshToken(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		"reset password"
// @Description	"reset password request"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.ResetPasswordRequest	true	"forgot password email payload"
// @Success	200		{string}	string						"OK"
// @Failure	404		{object}	dtos.ErrorResponse			"body"
// @Failure	500		{object}	dtos.ErrorResponse			"body"
// @Router		/auth/reset-password [post]
func (ctrl AuthController) ResetPassword(c *gin.Context) {
	var payload auth.ResetPasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.authsrv.ResetPassword(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		"forgot password"
// @Description	"forgot password request"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.ResendEmailRequest	true	"forgot password email payload"
// @Success	200		{string}	string					"OK"
// @Failure	404		{object}	dtos.ErrorResponse		"body"
// @Failure	500		{object}	dtos.ErrorResponse		"body"
// @Router		/auth/forgot-password [post]
func (ctrl AuthController) ForgotPassword(c *gin.Context) {
	var payload auth.ResendEmailRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.authsrv.ForgotPassword(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		"resend confirmation email"
// @Description	"Resend confirmation email to user"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.ResendEmailRequest	true	"resend confirmation email payload"
// @Success	200		{string}	string					"OK"
// @Failure	404		{object}	dtos.ErrorResponse		"body"
// @Failure	500		{object}	dtos.ErrorResponse		"body"
// @Router		/auth/resend-confirmation-email [post]
func (ctrl AuthController) ResendConfirmationEmail(c *gin.Context) {
	var payload auth.ResendEmailRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.authsrv.ResendConfirmEmail(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		"verify user email"
// @Description	"Verify a user email address"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Param		payload	body		auth.VerifyEmailRequest	true	"verify email payload"
// @Success	200		{string}	string					"OK"
// @Failure	400		{object}	dtos.ErrorResponse		"body"
// @Failure	404		{object}	dtos.ErrorResponse		"body"
// @Failure	500		{object}	dtos.ErrorResponse		"body"
// @Router		/auth/verify-email [post]
func (ctrl AuthController) VerifyEmail(c *gin.Context) {
	var payload auth.VerifyEmailRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		validator.GetErrors(c, err)
		return
	}

	err := ctrl.authsrv.VerifyEmail(&payload)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		"Google OAuth Login"
// @Description	"Login via Google OAuth"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Failure	400		{object}	dtos.ErrorResponse		"body"
// @Router		/auth/google-login [get]
func (ctrl AuthController) GoogleLogin(c *gin.Context) {
	ctrl.authsrv.GoogleOAuthLogin(c)
}

// @Summary		"Google OAuth Logout"
// @Description	"Logout via Google OAuth"
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Schemes
// @Security	BearerAuth
// @Failure	400		{object}	dtos.ErrorResponse		"body"
// @Router		/auth/google-logout [get]
func (ctrl AuthController) GoogleLogout(c *gin.Context) {
	ctrl.authsrv.GoogleOAuthLogout(c)
}

func (ctrl AuthController) GoogleLoginCallback(c *gin.Context) {
	response, err := ctrl.authsrv.GoogleOAuthCallbackLogin(c)
	if err != nil {
		body := err.(dtos.ErrorResponse)
		c.JSON(body.Status, body)
		return
	}

	c.JSON(http.StatusOK, response)
}
