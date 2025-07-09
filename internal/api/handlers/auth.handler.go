package handlers

import (
	"github.com/boldd/internal/application/auth"
	"github.com/boldd/internal/infrastructure/validator"
	"github.com/gin-gonic/gin"
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
}

func (ctrl AuthController) Login(c *gin.Context)                   {}
func (ctrl AuthController) Logout(c *gin.Context)                  {}
func (ctrl AuthController) ForgotPassword(c *gin.Context)          {}
func (ctrl AuthController) ResetPassword(c *gin.Context)           {}
func (ctrl AuthController) ResendConfirmationEmail(c *gin.Context) {}
func (ctrl AuthController) VerifyEmail(c *gin.Context)             {}
func (ctrl AuthController) GoogleLogin(c *gin.Context)             {}
