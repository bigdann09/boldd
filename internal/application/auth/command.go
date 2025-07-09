package auth

type IAuthCommandService interface {
	Login()
	Logout()
	Register()
	VerifyEmail()
	RefreshToken()
	ForgotPassword()
	ResendConfirmEmail()
}

type AuthCommandService struct {
}

func NewAuthCommandService() *AuthCommandService {
	return &AuthCommandService{}
}

func (srv *AuthCommandService) Register() {

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
