package sender

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendEmails(email Email) error {
  // Tries to connect with password
	fmt.Printf("INFO: Trying connection to %v on port %v\n", email.SmtpHost, email.SmtpPort)
	d := gomail.NewDialer(email.SmtpHost, email.SmtpPort, email.Username, email.Password)
	s, err := d.Dial()
	if err != nil {
		fmt.Println("ERROR: ", err)
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
			m.SetBody("text/html", email.Body)
		}

    // Tries to send the email and logs the result
		if err := gomail.Send(s, m); err != nil {
			fmt.Printf("ERROR: %v", err)
		}
		fmt.Printf("OK\n")

    // Resets the email and continues until every single email is sent
		m.Reset()
	}
	return nil
}
