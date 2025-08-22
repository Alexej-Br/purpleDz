// Package pkg
package pkg

import (
	"crypto/tls"
	"fmt"
	"go-adv/http/configs"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(conf *configs.Config, to string, subject string, text string) error {
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         conf.Email.Address,
	}
	from := fmt.Sprintf("My project <%s>", conf.Email.Email)
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	e.HTML = []byte("<h1>" + text + "</h1>")
	err := e.SendWithTLS("smtp.mail.ru:465",
		smtp.PlainAuth("", conf.Email.Email, conf.Email.Pass, conf.Email.Address),
		tlsconfig)
	if err != nil {
		return err
	}
	return nil
}
