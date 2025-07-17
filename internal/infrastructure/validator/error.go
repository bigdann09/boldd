package validator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// handle validator error
func validationErrors(err error) []ValidationError {
	var validationErrors []ValidationError
	fmt.Println(err)
	if err, ok := err.(validator.ValidationErrors); ok {
		for _, e := range err {
			validationErrors = append(validationErrors, ValidationError{
				Field: strings.ToLower(e.Field()),
				Error: customMessages(e),
			})
		}
	}
	return validationErrors
}

// customize validation error messages
func customMessages(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s field is required", strings.ToLower(fe.Field()))
	case "email":
		return "provided email is invalid"
	case "min":
		return fmt.Sprintf("%s minimum length is %s characters", strings.ToLower(fe.Field()), fe.Param())
	case "max":
		return fmt.Sprintf("%s maximum length is %s characters", strings.ToLower(fe.Field()), fe.Param())
	case "unique_email":
		return fmt.Sprintf("%s (%s) has already been taken", strings.ToLower(fe.Field()), fe.Value())
	}
	return "invalid validation error"
}

// build validation json response
func GetErrors(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors(err)})
}
