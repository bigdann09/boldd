package handlers

import (
	"github.com/boldd/internal/application/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authsrv auth.IAuthCommandService
}

func NewAuthController(authsrv auth.IAuthCommandService) *AuthController {
	return &AuthController{authsrv}
}

func (ctrl AuthController) Register(c *gin.Context)                {}
func (ctrl AuthController) Login(c *gin.Context)                   {}
func (ctrl AuthController) Logout(c *gin.Context)                  {}
func (ctrl AuthController) ForgotPassword(c *gin.Context)          {}
func (ctrl AuthController) ResetPassword(c *gin.Context)           {}
func (ctrl AuthController) ResendConfirmationEmail(c *gin.Context) {}
func (ctrl AuthController) VerifyEmail(c *gin.Context)             {}
func (ctrl AuthController) GoogleLogin(c *gin.Context)             {}
