package sender

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"

	"github.com/otaleghani/sbes/internal/oauth2"
)

func SendEmailOAuth(
	email Email,
) {
	d := gomail.NewDialer(email.SmtpHost, email.SmtpPort, email.Username, email.Password)
	d.TLSConfig = &tls.Config{ServerName: email.SmtpHost}
	d.Auth = oauth2.NewOauth2Authenticator(email.Username, email.Oauth)

	s, err := d.Dial()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
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
}
