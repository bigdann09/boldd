package mail

import (
	"embed"
	"log"
	"os"

	mailv2 "gopkg.in/mail.v2"
)

//go:embed "templates"
var templates embed.FS

type IMail interface {
	To(email string) *Mail
	Subject(subject string) *Mail
	Body(body ITemplate) *Mail
	Attachment(file *os.File) *Mail
	Send() error
	SendAsync() error
}

type Mail struct {
	from     string
	username string
	password string
	host     string
	port     int
	mailer   Mailer
}

type Mailer struct {
	to          string
	body        string
	subject     string
	attachments []os.File
}

type ITemplate interface {
	Path() string
	Template() string
}

func NewMail(from, username, password, host string, port int) *Mail {
	return &Mail{
		from:     from,
		username: username,
		password: password,
		host:     host,
		port:     port,
		mailer:   Mailer{attachments: make([]os.File, 5)},
	}
}

func (mail *Mail) To(email string) *Mail {
	log.Println("Assigning email to send to")
	mail.mailer.to = email
	return mail
}

func (mail *Mail) Subject(subject string) *Mail {
	log.Println("Setting up mail subject")
	mail.mailer.subject = subject
	return mail
}

func (mail *Mail) Body(body ITemplate) *Mail {
	log.Println("Bulding email body")
	mail.mailer.body = body.Template()
	return mail
}

func (mail *Mail) Attachment(file *os.File) *Mail {
	return mail
}

func (mail *Mail) Send() error {
	log.Println("sending mail to user")
	m := mailv2.NewMessage()
	m.SetHeader("From", mail.from)
	m.SetHeader("Subject", mail.mailer.subject)
	m.SetHeader("To", mail.mailer.to)
	m.SetBody("text/html", mail.mailer.body)

	dailer := mailv2.NewDialer("localhost", 1025, mail.username, mail.password)
	err := dailer.DialAndSend(m)
	return err
}

func (mail *Mail) SendAsync() error {
	return nil
}
