package auth

import (
	"github.com/boldd/internal/domain/user"
	"github.com/boldd/internal/infrastructure/auth/jwt"
	"github.com/boldd/pkgs/utils"
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
}

func NewAuthCommandService(userRepository user.IUserRepository, tokensrv jwt.ITokenService) *AuthCommandService {
	return &AuthCommandService{userRepository, tokensrv}
}

func (srv *AuthCommandService) Register(payload *RegisterRequest) (*AuthResponse, interface{}) {
	// register user
	newUser := user.NewUser(
		payload.FullName,
		payload.Email,
		payload.PhoneNumber,
		utils.HashPassword(payload.Password),
	)
	err := srv.userRepository.Create(newUser)
	if err != nil {
		return &AuthResponse{}, map[string]interface{}{"error": err, "code": 500}
	}

	// Assign Roles
	// err = srv.userRepository.AssignRole(newUser.ID, "customer")

	// TODO: send mail

	// return response payload
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
