package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID          string
	Fullname      string
	GoogleID      string
	Email         string
	PhoneNumber   string
	Password      string
	EmailVerified bool
}

func NewUser(fullname, email, phoneNumber, password string) *User {
	return &User{
		UUID:        uuid.NewString(),
		Fullname:    fullname,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
}

func NewGoogleUser(fullname, email string) *User {
	return &User{
		UUID:          uuid.NewString(),
		Fullname:      fullname,
		Email:         email,
		EmailVerified: true,
	}
}

func (u *User) GetFullName() string {
	return u.Fullname
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) VerifyEmail() {
	u.EmailVerified = true
}
