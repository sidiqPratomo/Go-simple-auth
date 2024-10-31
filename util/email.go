package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
)

type EmailHelper interface {
	AddRequest(to []string, subject string)
	CreateBody(emailTemplate string, data interface{}) error
	SendEmail() error
}

type emailHelperIpl struct {
	config       config.Config
	emailRequest emailRequest
}

func NewEmailHelperIpl(config *config.Config) emailHelperIpl {
	return emailHelperIpl{
		config: *config,
	}
}

type emailRequest struct {
	from    string
	to      []string
	subject string
	body    string
}

func (eh *emailHelperIpl) AddRequest(to []string, subject string) {
	eh.emailRequest = emailRequest{
		from:    eh.config.SendEmailIdentity,
		to:      to,
		subject: subject,
	}
}

func (eh *emailHelperIpl) SendEmail() error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + eh.emailRequest.subject + "!\n"
	msg := []byte(subject + mime + "\n" + eh.emailRequest.body)

	addr := fmt.Sprintf("%s:%s", eh.config.SendEmailHost, eh.config.SendEmailPort)

	auth := smtp.PlainAuth(eh.config.SendEmailIdentity, eh.config.SendEmailUsername, eh.config.SendEmailPassword, eh.config.SendEmailHost)

	if err := smtp.SendMail(addr, auth, eh.emailRequest.from, eh.emailRequest.to, msg); err != nil {
		return err
	}

	return nil
}

func (eh *emailHelperIpl) CreateBody(emailTemplate string, data interface{}) error {
	t, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		return err
	}

	eh.emailRequest.body = buf.String()

	return nil
}
