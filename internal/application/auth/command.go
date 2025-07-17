package auth

import (
	"github.com/boldd/internal/domain/user"
	"github.com/boldd/internal/infrastructure/auth/jwt"
	"github.com/boldd/pkgs/utils"
	"go.uber.org/zap"
)

type IAuthCommandService interface {
	Login()
	Logout()
	Register(payload *RegisterRequest) (*AuthResponse, interface{})
	VerifyEmail()
	RefreshToken()
	ForgotPassword()
	ResendConfirmEmail()
}

type AuthCommandService struct {
	userRepository user.IUserRepository
	tokensrv       jwt.ITokenService
	logger         *zap.Logger
}

func NewAuthCommandService(userRepository user.IUserRepository, tokensrv jwt.ITokenService, logger *zap.Logger) *AuthCommandService {
	return &AuthCommandService{userRepository, tokensrv, logger}
}

func (srv *AuthCommandService) Register(payload *RegisterRequest) (*AuthResponse, interface{}) {
	srv.logger.Info("adding new user to database")
	newUser := user.NewUser(
		payload.FullName,
		payload.Email,
		payload.PhoneNumber,
		utils.HashPassword(payload.Password),
	)
	err := srv.userRepository.Create(newUser)
	if err != nil {
		srv.logger.Error("there was an error adding user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": err, "code": 500}
	}

	srv.logger.Info("assign customer role to new user")
	if err = srv.userRepository.AssignRole(int(newUser.ID), "customer"); err != nil {
		srv.logger.Error("error assigning role to user", zap.Error(err))
		// TODO: Delete user record
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": 500}
	}

	// TODO: send mail

	return &AuthResponse{
		AccessToken:  srv.tokensrv.GenerateAccessToken(int(newUser.ID), "customer"),
		RefreshToken: srv.tokensrv.GenerateRefreshToken(int(newUser.ID)),
	}, nil
}

func (srv *AuthCommandService) Login() {

}

func (srv *AuthCommandService) RefreshToken() {

}

func (srv *AuthCommandService) Logout() {

}

func (srv *AuthCommandService) ForgotPassword() {

}

func (srv *AuthCommandService) ResetPassword() {

}

func (srv *AuthCommandService) VerifyEmail() {

}

func (srv *AuthCommandService) ResendConfirmEmail() {

}
