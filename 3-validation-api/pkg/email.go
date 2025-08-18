// Package pkg
package pkg

import (
	"fmt"
	"go-adv/http/configs"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(conf *configs.Config, to string, subject string, text string) {
	from := fmt.Sprintf("My project <%s>", conf.Email.Email)
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", conf.Email.Email, conf.Email.Pass, conf.Email.Address))
	if err != nil {
		log.Println(err)
	}
}
