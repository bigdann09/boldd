package mail

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type ResendConfirmationMail struct {
	filename string
	otp      int
}

func NewResendConfirmationMail(otp int) *RegistrationMail {
	return &RegistrationMail{
		filename: "resend.html",
		otp:      otp,
	}
}

func (tmpl ResendConfirmationMail) Path() string {
	path := path.Join("templates", tmpl.filename)
	return path
}

func (tmpl ResendConfirmationMail) FileExists() bool {
	_, err := os.Stat(tmpl.Path())
	return !errors.Is(err, os.ErrNotExist)
}

func (tmpl ResendConfirmationMail) Template() string {
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
