package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID          string
	Fullname      string
	Email         string
	PhoneNumber   string
	Password      string
	ReferralToken string
}

func NewUser(fullname, email, phoneNumber, password string) *User {
	return &User{
		UUID:          uuid.NewString(),
		Fullname:      fullname,
		Email:         email,
		Password:      password,
		PhoneNumber:   phoneNumber,
		ReferralToken: "",
	}
}

func (u *User) GetFullName() string {
	return u.Fullname
}

func (u *User) GetEmail() string {
	return u.Email
}
