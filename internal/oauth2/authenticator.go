package oauth2

import (
	"fmt"
	"net/smtp"
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

func NewOauth2Authenticator(username, accessToken string) smtp.Auth {
	return &oauth2Authenticator{
		username:    username,
		accessToken: accessToken,
	}
}
