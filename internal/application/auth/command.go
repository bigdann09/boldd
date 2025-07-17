package auth

import (
	"github.com/boldd/internal/domain/user"
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
}

func NewAuthCommandService(userRepository user.IUserRepository) *AuthCommandService {
	return &AuthCommandService{userRepository}
}

func (srv *AuthCommandService) Register(payload *RegisterRequest) (*AuthResponse, interface{}) {
	// register user
	newUser := user.NewUser(payload.FullName, payload.Email, payload.PhoneNumber, utils.HashPassword(payload.Password))
	err := srv.userRepository.Create(newUser)
	if err != nil {
		return &AuthResponse{}, map[string]interface{}{"error": err, "code": 500}
	}

	// TODO: send mail

	// return response payload
	return &AuthResponse{
		AccessToken:  "access token",
		RefreshToken: "refresh token",
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
