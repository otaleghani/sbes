package sender

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func TestSMTPConnection(smtpHost string, smtpPort int, username, password string) error {
	d := gomail.NewDialer(smtpHost, smtpPort, username, password)

	// Try to establish a connection
	c, err := d.Dial()
	if err != nil {
		return err
	}
	defer c.Close()
	return nil
}

// func sendtestemail

func SendEmails(smtpHost string, smtpPort int, username, password, from string, to []string, subject, body string) error {
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

// func main() {
//     // Define your SMTP server details
//     smtpHost := "smtp.gmail.com"
//     smtpPort := 587
//
//     // Gmail account credentials
//     username := "o.taleghani@talesign.com"
//     password := ""
//
//     // Email details
//     from := "o.taleghani@talesign.com"
//     to := []string{"account@aficleaning.com", "o.taleghani@talesign.com"} // List of recipients
//     subject := "Test Email"
//     body := "<h1>Hello, this is a test email!</h1>"
//
// }
