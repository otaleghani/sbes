package sender

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
  "fmt"
  
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

	m := gomail.NewMessage()
  for _, recipient := range email.Mailing_List {
	  m.SetHeader("From", email.Username)
	  m.SetHeader("To", recipient)
	  m.SetHeader("Subject", email.Subject)
    if email.Msg_Type == ".txt" {
	    m.SetBody("text/plain", email.Body)
    } else {
	    m.SetBody("text/html", email.Body)
    }
    if err := gomail.Send(s, m); err != nil {
		  fmt.Printf("Could not send email to %v: %v", recipient, err)
	  }
  }
  m.Reset()
}
