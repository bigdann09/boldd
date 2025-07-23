package dtos

type User struct {
	ID            uint   `json:"-"`
	Email         string `json:"email"`
	UUID          string `json:"uuid"`
	Fullname      string `json:"fullname"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `json:"-"`
	EmailVerified bool   `json:"email_verified"`
}

type UserResponse struct {
	User
	Roles []string `json:"roles"`
}

func (u *User) Convert(roles ...string) UserResponse {
	var response UserResponse
	response.ID = u.ID
	response.UUID = u.UUID
	response.Roles = roles
	response.Email = u.Email
	response.Fullname = u.Fullname
	response.Password = u.Password
	response.PhoneNumber = u.PhoneNumber
	response.EmailVerified = u.EmailVerified
	return response
}
