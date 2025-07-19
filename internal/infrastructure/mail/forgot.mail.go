package mail

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type ForgotPasswordMail struct {
	filename string
	otp      int
}

func NewForgotPasswordMail(otp int) *RegistrationMail {
	return &RegistrationMail{
		filename: "reset.html",
		otp:      otp,
	}
}

func (tmpl ForgotPasswordMail) Path() string {
	path := path.Join("templates", tmpl.filename)
	return path
}

func (tmpl ForgotPasswordMail) FileExists() bool {
	_, err := os.Stat(tmpl.Path())
	return !errors.Is(err, os.ErrNotExist)
}

func (tmpl ForgotPasswordMail) Template() string {
	// TODO: validate file
	fmt.Println("file exists", tmpl.FileExists())
	var body strings.Builder
	template, err := template.ParseFS(templates, tmpl.Path())
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	template.Execute(&body, map[string]interface{}{"otp": tmpl.otp})
	return body.String()
}
