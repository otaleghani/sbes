package sender

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
	"github.com/otaleghani/sbes/internal/oauth2"
	"github.com/otaleghani/sbes/internal/tracker"
)

func SendEmailOAuth(email Email) {
	// Tries to connect with OAuth2 auth
	fmt.Printf("INFO: Trying connection to %v on port %v\n", email.SmtpHost, email.SmtpPort)
	d := gomail.NewDialer(email.SmtpHost, email.SmtpPort, email.Username, email.Password)
	d.TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: email.SmtpHost,
	}
	d.Auth = oauth2.NewOauth2Authenticator(email.Username, email.Oauth)
	s, err := d.Dial()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("INFO: Connection to host successful. Preparing emails...")

	// Creates a new message
	m := gomail.NewMessage()

	// Cycles through the mailing list
	for _, recipient := range email.MailingList {
		fmt.Printf("LOG: Sending email to %-42v Result: ", recipient)

		// Sets the header for the email
		m.SetHeader("From", email.Username)
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", email.Subject)
		if email.MsgType == ".txt" {
			m.SetBody("text/plain", email.Body)
		} else {
      body := tracker.AddTrackerToEmail(email.Domain, recipient, email.Campaign, email.Body)
			m.SetBody("text/html", body)
		}

		// Tries to send the email and logs the result
		if err := gomail.Send(s, m); err != nil {
			fmt.Printf("ERROR: %v", err)
		}
		fmt.Printf("OK\n")

		// Resets the email and continues until every single email is sent
		m.Reset()
	}
}
