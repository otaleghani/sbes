package oauth2

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

type oauth2Authenticator struct {
    username, accessToken string
}

func (a *oauth2Authenticator) Start(server *smtp.ServerInfo) (string, []byte, error) {
    authString := fmt.Sprintf("user=%s\x01auth=Bearer %s\x01\x01", a.username, a.accessToken)
    return "XOAUTH2", []byte(authString), nil
}

func (a *oauth2Authenticator) Next(fromServer []byte, more bool) ([]byte, error) {
    if more {
        return nil, fmt.Errorf("unexpected server challenge: %s", fromServer)
    }
    return nil, nil
}

func newOauth2Authenticator(username, accessToken string) smtp.Auth {
	return &oauth2Authenticator{
		username:    username,
		accessToken: accessToken,
	}
}

// func sendEmail(accessToken string) {
// 	d := gomail.NewDialer("smtp.gmail.com", 587, "o.taleghani@talesign.com", "")
// 	d.TLSConfig = &tls.Config{ServerName: "smtp.gmail.com"}
// 
// 	d.Auth = newOauth2Authenticator("o.taleghani@talesign.com", accessToken)
// 
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "o.taleghani@talesign.com")
// 	m.SetHeader("To", "o.taleghani@talesign.com")
// 	m.SetHeader("Subject", "Hello")
// 	m.SetBody("text/plain", "Hello from gomail with OAuth2!")
// 
// 	if err := d.DialAndSend(m); err != nil {
// 		fmt.Println(err)
// 	}
// }
