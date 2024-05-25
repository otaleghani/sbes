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

func SendEmails(email Email) error {

	d := gomail.NewDialer(email.SmtpHost, email.SmtpPort, email.Username, email.Password)
	s, err := d.Dial()
	if err != nil {
    fmt.Println("ERROR: ", err)
	}

  fmt.Println("INFO: Connection to host successful. Preparing emails...")

	m := gomail.NewMessage()
  for _, recipient := range email.Mailing_List {
    fmt.Printf("LOG: Sending email to %v. Result: ", recipient)
		m.SetHeader("From", email.Username)
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", email.Subject)
    if email.Msg_Type == ".txt" {
	    m.SetBody("text/plain", email.Body)
    } else {
	    m.SetBody("text/html", email.Body)
    }

		if err := gomail.Send(s, m); err != nil {
      fmt.Printf("ERROR: %v", err)
		}
		fmt.Printf("OK\n")
		m.Reset()
	}
	return nil
}
