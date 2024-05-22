package main

import (
	"fmt"
	"strings"
  "strconv"
  "github.com/otaleghani/sbes/internal/terminalinput"
)

type Config struct {
	Username string `username`
	Password string `password`
	SmtpHost string `json:smtpHost`
	SmtpPort int `json:smtpPort`
}

type Database struct {
	Accounts map[string]Config `json:account`
}

func handleAuthAdd() error {
  user := strings.TrimSpace(
    terminalinput.ReadInput("\nEnter username\n\r-> "))

  pass := strings.TrimSpace(
    terminalinput.ReadInput("\nEnter password\n\r-> "))

  host := strings.TrimSpace(
    terminalinput.ReadInput("\nEnter host\n\r-> "))

  port, err := strconv.Atoi(
    strings.TrimSpace(
      terminalinput.ReadInput("\nEnter port\n\r-> ")))
	if err != nil {
		return err
	}

	fmt.Printf("\nTesting connection...\n\tUsername: %v,\n\tPassword: %v\n\tHost: %v,\n\tPort: %v\n", user, pass, host, port)

  err = TestSMTPConnection(host, port, user, pass)
	if err != nil {
    return err
	}

  fmt.Println("Connection extablished. Saving data in local cache.")

  err = AuthAdd(
    Config{Username:user, Password:pass, SmtpHost:host, SmtpPort:port})
  if err != nil {
    return err
  }

  fmt.Println("Credential successfully added.")
	return nil
}

