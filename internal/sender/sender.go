package sender

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type Email struct {
  // Authentication
  SmtpHost string
  SmtpPort int
  Username string
  Password string
  Oauth string

  // Message
  From string
  Mailing_List []string
  Subject string
  Body string
  Msg_Type string
}

func SendEmails(
  smtpHost string,
  smtpPort int,
  username, password, from string,
  to []string,
  subject,
  body string,
) error {

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)
	s, err := d.Dial()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	for i := 0; i < len(to); i++ {
		m.SetHeader("From", from)
		m.SetHeader("To", to[i])
		m.SetHeader("Subject", subject)
		m.SetBody("text/plain", body)
		//m.SetBody("text/html", body)

		if err := gomail.Send(s, m); err != nil {
			fmt.Printf("Could not send email to %q: %v", to[i], err)
		}
		m.Reset()
	}
	return nil
}
