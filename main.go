package main

import (
	// "fmt"
	// "log"
	"gopkg.in/gomail.v2"
	// "os"
)

func sendEmail(smtpHost string, smtpPort int, username, password, from string, to []string, subject, body string) error {
	for i := 0; i < len(to); i++ {
		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", to[i])
		m.SetHeader("Subject", subject)
		m.SetBody("text/plain", body)
		//m.SetBody("text/html", body)

		d := gomail.NewDialer(smtpHost, smtpPort, username, password)

		// Send the email
		if err := d.DialAndSend(m); err != nil {
			return err
		}
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
//     password := "xvgy yxjh gpkm idjn"
//
//     // Email details
//     from := "o.taleghani@talesign.com"
//     to := []string{"account@aficleaning.com", "o.taleghani@talesign.com"} // List of recipients
//     subject := "Test Email"
//     body := "<h1>Hello, this is a test email!</h1>"
//
//     err := sendEmail(smtpHost, smtpPort, username, password, from, to, subject, body)
//     if err != nil {
//         log.Fatalf("Could not send email: %v", err)
//     }
//
//     fmt.Println("Email sent successfully!")
// }

func main() {
  repl()
}
