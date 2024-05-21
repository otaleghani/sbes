package main

import (
	"flag"
	"fmt"
  "os"
)

func repl() {


	switch os.Args[1] {
	case "send":
    sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	  recipients := sendCmd.String("r", "Recipients list", "Path to the recipients .csv file.")
	  auth := sendCmd.String("a", "Default Body", "Credentials added by using \"sbes auth\"")
	  message := sendCmd.String("m", "Default Body", "Message to be send. Either .txt for plaintext or .html")
    sendCmd.Parse(os.Args[2:])
    cmdSend(*recipients, *auth, *message)
	case "auth":
    authCmd := flag.NewFlagSet("auth", flag.ExitOnError)
    username := authCmd.String("u", "Username", "Username in complete address format")
    password := authCmd.String("p", "Password", "Parrword for the given account")
    authCmd.Parse(os.Args[2:])
    cmdAuth(*username, *password)
	case "help":
    cmdHelp()
	default:
    cmdHelp()
	}
}

func cmdHelp() {
  fmt.Println("Help ")
}


func cmdSend(recipients, auth, message string) {
  fmt.Println("Recipients: ", recipients)
  fmt.Println("Auth: ", auth)
  fmt.Println("Message: ", message)
}

func cmdAuth(username, password string) {
  // here you will need to add those
  fmt.Println("Auth")
  err := AddAuth(username, password)
  if err != nil {
    fmt.Println(err)
  }
}
