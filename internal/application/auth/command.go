package auth

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
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
	VerifyEmail(payload *VerifyEmailRequest) interface{}
	RefreshToken(payload *RefreshTokenRequest) (*AuthResponse, interface{})
	ResetPassword(payload *ResetPasswordRequest) interface{}
	ForgotPassword(payload *ResendEmailRequest) interface{}
	ResendConfirmEmail(payload *ResendEmailRequest) interface{}
}

type AuthCommandService struct {
	userRepository    repositories.IUserRepository
	otpRepository     repositories.IOtpRepository
	addressRepository repositories.IUserAddressRepository
	tokensrv          jwt.ITokenService
	logger            *zap.Logger
	mailer            mail.IMail
}

func NewAuthCommandService(
	userRepository repositories.IUserRepository,
	otpRepository repositories.IOtpRepository,
	addressRepository repositories.IUserAddressRepository,
	tokensrv jwt.ITokenService,
	logger *zap.Logger,
	mailer mail.IMail,
) *AuthCommandService {
	return &AuthCommandService{userRepository, otpRepository, addressRepository, tokensrv, logger, mailer}
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
		return &AuthResponse{}, map[string]interface{}{"error": err, "code": http.StatusInternalServerError}
	}

	srv.logger.Info("storing user addresses")
	for _, address := range payload.Addresses {
		srv.addressRepository.Create(entities.NewUserAddress(
			newUser.ID,
			address.State,
			address.City,
			address.ZipCode,
			address.Address,
		))
	}

	srv.logger.Info("assign customer role to new user")
	if err = srv.userRepository.AssignRole(int(newUser.ID), "customer"); err != nil {
		srv.logger.Info("deleting new user record")
		srv.userRepository.Delete(int(newUser.ID))
		srv.logger.Error("error assigning role to user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": http.StatusInternalServerError}
	}

	otpCode := utils.GenerateOTP()
	srv.logger.Info("store user otp for email verification")
	if err := srv.otpRepository.Create(entities.NewOtp(newUser.Email, otpCode, time.Now().Add(time.Minute*5))); err != nil {
		srv.logger.Error("error storing email code for user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": http.StatusInternalServerError}
	}

	srv.logger.Info("sending registration email to user")
	if err := srv.mailer.To(newUser.Email).Subject("Registration Complete").
		Body(mail.NewRegistrationMail(newUser.Fullname, otpCode)).
		Send(); err != nil {
		srv.logger.Error("error sending email to user", zap.Error(err))
		return &AuthResponse{}, map[string]interface{}{"error": "there was an error creating user account", "code": http.StatusInternalServerError}
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
		return &AuthResponse{}, map[string]interface{}{"error": "could not retrieve user data", "code": http.StatusInternalServerError}
	}

	srv.logger.Info("checking if record exists")
	if reflect.DeepEqual(user, &dtos.UserResponse{}) {
		srv.logger.Warn("user record not found", zap.String("email", payload.Email))
		return &AuthResponse{}, map[string]interface{}{"error": "invalid login credentials", "code": http.StatusBadRequest}
	}

	srv.logger.Info("validate password if email matches or record was found")
	if err := utils.ComparePasswords(user.Password, payload.Password); err != nil {
		srv.logger.Warn("user password was incorrect", zap.String("email", payload.Email))
		return &AuthResponse{}, map[string]interface{}{"error": "invalid login credentials", "code": http.StatusBadRequest}
	}

	srv.logger.Info("check if email is verified")
	if !user.EmailVerified {
		srv.logger.Warn("user email not verified", zap.String("email", payload.Email))
		return &AuthResponse{}, map[string]interface{}{"error": "email address not verified", "code": http.StatusBadRequest}
	}

	roles, err := srv.userRepository.Roles(user.ID)
	if err != nil {
		srv.logger.Error("could not retrieve user roles")
		return &AuthResponse{}, map[string]interface{}{"error": "error authenticating user", "code": http.StatusInternalServerError}
	}

	return &AuthResponse{
		AccessToken:  srv.tokensrv.GenerateAccessToken(int(user.ID), roles...),
		RefreshToken: srv.tokensrv.GenerateRefreshToken(int(user.ID)),
	}, nil
}

func (srv *AuthCommandService) RefreshToken(payload *RefreshTokenRequest) (*AuthResponse, interface{}) {
	// TODO: check if refresh token is still valid
	claims, err := srv.tokensrv.ValidateToken(payload.RefreshToken)
	if err != nil {
		return &AuthResponse{}, map[string]interface{}{"error": "refresh token expired", "code": http.StatusInternalServerError}
	}

	srv.logger.Info("retrieve user roles from claims")
	roles, _ := srv.userRepository.Roles(uint(claims.Id))

	return &AuthResponse{
		AccessToken:  srv.tokensrv.GenerateAccessToken(claims.Id, roles...),
		RefreshToken: payload.RefreshToken,
	}, nil
}

func (srv *AuthCommandService) Logout() {

}

func (srv *AuthCommandService) ForgotPassword(payload *ResendEmailRequest) interface{} {
	srv.logger.Info("checking if email address is registered")
	if exists := srv.userRepository.EmailExists(payload.Email); !exists {
		srv.logger.Warn("user record not found", zap.String("email", payload.Email))
		return map[string]interface{}{"error": "invalid email address", "code": http.StatusNotFound}
	}

	srv.logger.Info("check if otp ecord exists for email", zap.String("email", payload.Email))
	if exists := srv.otpRepository.Exists(payload.Email); exists {
		srv.logger.Info("deleting old otp record")
		srv.otpRepository.DeleteByEmail(payload.Email)
	}

	otpCode := utils.GenerateOTP()
	srv.logger.Info("store user otp for email verification")
	err := srv.otpRepository.Create(entities.NewOtp(payload.Email, otpCode, time.Now().Add(time.Minute*5)))
	if err != nil {
		srv.logger.Error("error storing email code for user", zap.Error(err))
		return map[string]interface{}{"error": "could not send reset email", "code": http.StatusInternalServerError}
	}

	err = srv.mailer.To(payload.Email).Subject("Reset password").
		Body(mail.NewForgotPasswordMail(otpCode)).Send()
	if err != nil {
		return map[string]interface{}{"error": "could not send reset email", "code": http.StatusInternalServerError}
	}

	return nil
}

func (srv *AuthCommandService) ResetPassword(payload *ResetPasswordRequest) interface{} {
	srv.logger.Info("retrieving user information")
	user, _ := srv.userRepository.FindByEmail(payload.Email)

	srv.logger.Info("checking if record exists")
	if reflect.DeepEqual(user, &dtos.UserResponse{}) {
		srv.logger.Warn("user record not found", zap.String("email", payload.Email))
		return map[string]interface{}{"error": "record not found", "code": http.StatusNotFound}
	}

	srv.logger.Info("retrieve otp data")
	response, err := srv.otpRepository.Find(payload.Email)
	if err != nil {
		srv.logger.Error("encountered an error retrieiving otp info", zap.Error(err))
		return map[string]interface{}{"error": "there was an error verifying email", "code": http.StatusBadRequest}
	}

	code := strconv.Itoa(response.Code)
	if reflect.DeepEqual(response, &entities.Otp{}) || time.Now().After(response.ExpiresAt) || payload.Code != code {
		if strings.EqualFold(response.Email, "") {
			srv.logger.Info("deleting otp info if it exists..")
			srv.otpRepository.Delete(response.UUID)
		}
		srv.logger.Info("otp code expired for", zap.String("email", payload.Email))
		return map[string]interface{}{"error": "otp code expired or invalid", "code": http.StatusBadRequest}
	}

	srv.logger.Info("verifing email address", zap.String("email", payload.Email))
	if err = srv.userRepository.Update(user.ID, &entities.User{Password: utils.HashPassword(payload.Password)}); err != nil {
		srv.logger.Error("encountered an error updating email status", zap.Error(err))
		return map[string]interface{}{"error": "there was an error verifying email", "code": http.StatusInternalServerError}
	}

	return nil
}

func (srv *AuthCommandService) VerifyEmail(payload *VerifyEmailRequest) interface{} {
	srv.logger.Info("retrieving user information")
	user, _ := srv.userRepository.FindByEmail(payload.Email)

	srv.logger.Info("checking if record exists")
	if reflect.DeepEqual(user, &dtos.UserResponse{}) {
		srv.logger.Warn("user record not found", zap.String("email", payload.Email))
		return map[string]interface{}{"error": "record not found", "code": http.StatusNotFound}
	}

	srv.logger.Info("retrieve otp data")
	response, err := srv.otpRepository.Find(payload.Email)
	if err != nil {
		srv.logger.Error("encountered an error retrieiving otp info", zap.Error(err))
		return map[string]interface{}{"error": "there was an error verifying email", "code": http.StatusBadRequest}
	}

	code := strconv.Itoa(response.Code)
	if reflect.DeepEqual(response, &entities.Otp{}) || time.Now().After(response.ExpiresAt) || payload.Code != code {
		if strings.EqualFold(response.Email, "") {
			srv.logger.Info("deleting otp info if it exists..")
			srv.otpRepository.Delete(response.UUID)
		}
		srv.logger.Info("otp code expired for", zap.String("email", payload.Email))
		return map[string]interface{}{"error": "otp code expired or invalid", "code": http.StatusBadRequest}
	}

	srv.logger.Info("delete otp record after verification")
	if err = srv.otpRepository.Delete(response.UUID); err != nil {
		srv.logger.Error("encountered an error deleting otp record", zap.Error(err))
		return map[string]interface{}{"error": "there was an error verifying email", "code": http.StatusInternalServerError}
	}

	srv.logger.Info("verifing email address", zap.String("email", payload.Email))
	if err = srv.userRepository.Update(user.ID, &entities.User{EmailVerified: true}); err != nil {
		srv.logger.Error("encountered an error updating email status", zap.Error(err))
		return map[string]interface{}{"error": "there was an error verifying email", "code": http.StatusInternalServerError}
	}

	return nil
}

func (srv *AuthCommandService) ResendConfirmEmail(payload *ResendEmailRequest) interface{} {
	srv.logger.Info("retrieving user information")
	user, _ := srv.userRepository.FindByEmail(payload.Email)

	srv.logger.Info("checking if record exists")
	if reflect.DeepEqual(user, &dtos.UserResponse{}) {
		srv.logger.Warn("user record not found", zap.String("email", payload.Email))
		return map[string]interface{}{"error": "invalid email address", "code": http.StatusNotFound}
	}

	if user.EmailVerified {
		return map[string]interface{}{"error": "email address already verified", "code": http.StatusBadRequest}
	}

	srv.logger.Info("check if otp ecord exists for email", zap.String("email", payload.Email))
	if exists := srv.otpRepository.Exists(payload.Email); exists {
		srv.otpRepository.DeleteByEmail(payload.Email)
	}

	otpCode := utils.GenerateOTP()
	srv.logger.Info("store user otp for email verification")
	err := srv.otpRepository.Create(entities.NewOtp(payload.Email, otpCode, time.Now().Add(time.Minute*5)))
	if err != nil {
		srv.logger.Error("error storing email code for user", zap.Error(err))
		return map[string]interface{}{"error": "could not resend confirmation email", "code": http.StatusInternalServerError}
	}

	err = srv.mailer.To(payload.Email).Subject("Verify Email Address").
		Body(mail.NewResendConfirmationMail(otpCode)).Send()
	if err != nil {
		return map[string]interface{}{"error": "could not resend confirmation email", "code": http.StatusInternalServerError}
	}

	return nil
}
