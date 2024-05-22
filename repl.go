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

	case "auth-add":
		authCmd := flag.NewFlagSet("auth", flag.ExitOnError)
		username := authCmd.String("u", "Username", "Username in complete address format")
		password := authCmd.String("p", "Password", "Parrword for the given account")
		authCmd.Parse(os.Args[2:])
		cmdAuthAdd(*username, *password)

	case "auth-delete":
		authCmd := flag.NewFlagSet("auth", flag.ExitOnError)
		username := authCmd.String("u", "Username", "Username in complete address format")
		// password := authCmd.String("p", "Password", "Parrword for the given account")
		authCmd.Parse(os.Args[2:])
		cmdAuthDelete(*username)
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

func cmdAuthAdd(username, password string) {
	// fmt.Println("Auth add")
	// err := AddAuth(username, password)
	err := handleAuthAdd()
	if err != nil {
		fmt.Println(err)
	}
}

func cmdAuthDelete(username string) {
	fmt.Println("Auth delete")
	// err := DeleteAuth(username)
	// if err != nil {
	//   fmt.Println(err)
	// }
}
