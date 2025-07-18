package dtos

type UserResponse struct {
	ID            uint   `json:"_"`
	Email         string `json:"email"`
	UUID          string `json:"uuid"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `json:"_"`
	EmailVerified bool   `json:"email_verified"`
}
