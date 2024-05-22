package main

import (
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

