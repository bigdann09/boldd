package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	validator := &Validator{
		validator: v,
	}
	return validator
}

func (v *Validator) RegisterValidators() {

}
