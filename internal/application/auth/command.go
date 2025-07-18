package auth

import (
	"reflect"
	"time"

	"github.com/boldd/internal/domain/dtos"
	"github.com/boldd/internal/domain/entities"
	"github.com/boldd/internal/infrastructure/auth/jwt"
	"github.com/boldd/internal/infrastructure/mail"
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/boldd/pkgs/utils"
	"go.uber.org/zap"
)

type IAuthCommandService interface {
	Logout()
	Login(payload *LoginRequest) (*AuthResponse, interface{})
	Register(payload *RegisterRequest) (*AuthResponse, interface{})
	VerifyEmail()
	RefreshToken()
	ForgotPassword()
	ResendConfirmEmail()
}

type AuthCommandService struct {
	userRepository repositories.IUserRepository
	otpRepository  repositories.IOtpRepository
	tokensrv       jwt.ITokenService
	logger         *zap.Logger
	mailer         mail.IMail
}

func NewAuthCommandService(
	userRepository repositories.IUserRepository,
	otpRepository repositories.IOtpRepository,
	tokensrv jwt.ITokenService,
	logger *zap.Logger,
	mailer mail.IMail,
) *AuthCommandService {
	return &AuthCommandService{userRepository, otpRepository, tokensrv, logger, mailer}
}

func (srv *AuthCommandService) Register(payload *RegisterRequest) (*AuthResponse, interface{}) {
	srv.logger.Info("adding new user to database")
	newUser := entities.NewUser(
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
		srv.logger.Info("deleting new user record")
		srv.userRepository.Delete(int(newUser.ID))
		srv.logger.Error("error assigning role to user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": 500}
	}

	otpCode := utils.GenerateOTP()
	srv.logger.Info("store user otp for email verification")
	if err := srv.otpRepository.Create(entities.NewOtp(newUser.Email, otpCode, time.Now().Add(time.Minute*5))); err != nil {
		srv.logger.Error("error storing email code for user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": 500}
	}

	srv.logger.Info("sending registration email to user")
	if err := srv.mailer.To(newUser.Email).Subject("Registration Complete").
		Body(mail.NewRegistrationMail(newUser.Fullname, otpCode)).
		Send(); err != nil {
		srv.logger.Error("error sending email to user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": 500}
	}

	return &AuthResponse{
		AccessToken:  srv.tokensrv.GenerateAccessToken(int(newUser.ID), "customer"),
		RefreshToken: srv.tokensrv.GenerateRefreshToken(int(newUser.ID)),
	}, nil
}

func (srv *AuthCommandService) Login(payload *LoginRequest) (*AuthResponse, interface{}) {
	srv.logger.Info("authenticating user record")
	user, err := srv.userRepository.FindByEmail(payload.Email)
	if err != nil {
		srv.logger.Error("encountered an error retrieve user information", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "could not retrieve user data", "code": 500}
	}

	srv.logger.Info("checking if record exists")
	if reflect.DeepEqual(user, &dtos.UserResponse{}) {
		srv.logger.Warn("user record not found", zap.String("email", payload.Email))
		return &AuthResponse{}, map[string]interface{}{"error": "invalid login credentials", "code": 400}
	}

	srv.logger.Info("validate password if email matches or record was found")
	if err := utils.ComparePasswords(user.Password, payload.Password); err != nil {
		srv.logger.Warn("user password was incorrect", zap.String("email", payload.Email))
		return &AuthResponse{}, map[string]interface{}{"error": "invalid login credentials", "code": 400}
	}

	srv.logger.Info("check if email is verified")
	if !user.EmailVerified {
		srv.logger.Warn("user email not verified", zap.String("email", payload.Email))
		return &AuthResponse{}, map[string]interface{}{"error": "email address not verified", "code": 400}
	}

	roles, err := srv.userRepository.Roles(user.ID)
	if err != nil {
		srv.logger.Error("could not retrieve user roles")
		return &AuthResponse{}, map[string]interface{}{"error": "error authenticating user", "code": 500}
	}

	return &AuthResponse{
		AccessToken:  srv.tokensrv.GenerateAccessToken(int(user.ID), roles...),
		RefreshToken: srv.tokensrv.GenerateRefreshToken(int(user.ID)),
	}, nil
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
