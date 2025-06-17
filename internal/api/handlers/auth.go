package handlers

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ctrl AuthController) Register(c *gin.Context)                {}
func (ctrl AuthController) Login(c *gin.Context)                   {}
func (ctrl AuthController) Logout(c *gin.Context)                  {}
func (ctrl AuthController) ForgotPassword(c *gin.Context)          {}
func (ctrl AuthController) ResetPassword(c *gin.Context)           {}
func (ctrl AuthController) ResendConfirmationEmail(c *gin.Context) {}
func (ctrl AuthController) VerifyEmail(c *gin.Context)             {}
func (ctrl AuthController) GoogleLogin(c *gin.Context)             {}
