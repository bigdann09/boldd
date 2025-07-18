package mail

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type RegistrationMail struct {
	filename string
	otp      int
	fullname string
}

func NewRegistrationMail(fullname string, otp int) *RegistrationMail {
	return &RegistrationMail{
		filename: "register.html",
		otp:      otp,
		fullname: fullname,
	}
}

func (tmpl RegistrationMail) Path() string {
	path := path.Join("templates", tmpl.filename)
	return path
}

func (tmpl RegistrationMail) FileExists() bool {
	_, err := os.Stat(tmpl.Path())
	return !errors.Is(err, os.ErrNotExist)
}

func (tmpl RegistrationMail) Template() string {
	// TODO: validate file
	fmt.Println("file exists", tmpl.FileExists())
	var body strings.Builder
	template, err := template.ParseFS(templates, tmpl.Path())
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	template.Execute(&body, map[string]interface{}{"otp": tmpl.otp, "fullname": tmpl.fullname})
	return body.String()
}
