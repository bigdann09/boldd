package validator

import (
	"github.com/boldd/internal/infrastructure/persistence/repositories"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Validator struct {
	validator      *validator.Validate
	userRepository repositories.IUserRepository
}

func NewValidator(db *gorm.DB) *Validator {
	v := validator.New()
	validator := &Validator{
		validator:      v,
		userRepository: repositories.NewUserRepository(db),
	}
	return validator
}

func (v *Validator) RegisterValidators() {
	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("unique_email", v.uniqueEmail)
	}
}

func (v *Validator) uniqueEmail(fl validator.FieldLevel) bool {
	email := fl.Field()

	itExists := v.userRepository.EmailExists(email.String())
	return !itExists
}
