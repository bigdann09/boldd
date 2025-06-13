package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func NewUser(firstname, lastname, email, password string) *User {
	return &User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
	}
}
func (u *User) GetFullName() string {
	return u.Firstname + " " + u.Lastname
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) UpdateUser(firstname, lastname, email, password string) {
	u.Firstname = firstname
	u.Lastname = lastname
	u.Email = email
	u.Password = password
}
