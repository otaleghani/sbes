package main

import (
	"fmt"
	"strings"
  "strconv"
  "github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdAuthAdd() {
	err := handleAuthAdd()
	if err != nil {
    fmt.Println("\nERROR:", err)
	}
}

func cmdAuthDelete(username string) {
  err := handleAuthDelete(username)
  if err != nil {
    fmt.Println("\nERROR:", err)
  }
}

func handleAuthAdd() error {
  user := strings.TrimSpace(
    terminalinput.ReadInput("Enter username\n\r-> "))

  pass := strings.TrimSpace(
    terminalinput.ReadInput("Enter password\n\r-> "))

  host := strings.TrimSpace(
    terminalinput.ReadInput("Enter host\n\r-> "))

  port, err := strconv.Atoi(
    strings.TrimSpace(
      terminalinput.ReadInput("Enter port\n\r-> ")))
	if err != nil {
		return err
	}

	fmt.Printf("\nTesting connection...\n\tUsername: %v,\n\tPassword: %v\n\tHost: %v,\n\tPort: %v\n\n", user, pass, host, port)

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

func handleAuthDelete(user string) error {
  // We do not care if it was there or not...
  accept := strings.TrimSpace(
    terminalinput.ReadInput("\033[031m\nAre you sure (y/N) \n\r-> "))

  if accept == "n" {
    return nil
  } 

  err := AuthDelete(user)
  if err != nil {
    return err
  }
  return nil
}

