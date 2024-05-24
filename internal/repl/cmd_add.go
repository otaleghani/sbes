package repl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/oauth2"
	"github.com/otaleghani/sbes/internal/parser"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdAdd_Account() {
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
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("\nTesting connection...\n\tUsername: %v,\n\tPassword: %v\n\tHost: %v,\n\tPort: %v\n\n", user, pass, host, port)
	if err := database.AccountAdd(user, pass, host, port); err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("SUCCESS: Account saved")
}

func cmdAdd_OAuthClient() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter a name for this OAuth Client\n\r-> "))
	id := strings.TrimSpace(
		terminalinput.ReadInput("Enter OAuth Client Id\n\r-> "))
	secret := strings.TrimSpace(
		terminalinput.ReadInput("Enter OAuth Client Secret\n\r-> "))
	// To do: add test for determining if the id/secret combination is good
	if err := database.OAuthClientAdd(name, id, secret); err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("SUCCESS: OAuth Client saved")
}

func cmdAdd_OAuthToken() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter a name for this OAuth Token\n\r-> "))
	nameClient := strings.TrimSpace(
		terminalinput.ReadInput("Enter the name of the Client that you want to use\n\r-> "))
	id, secret, err := database.OAuthClientFind(nameClient)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	token := oauth2.GetOauth2(id, secret)
	if token == "" {
		fmt.Println("ERROR: Token not generated. Check you client credentials.")
		return
	}
	if err := database.OAuthTokenAdd(name, token); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: OAuth Client saved")
}

func cmdAdd_MailingList() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name for mailing list\n\r-> "))
	filePath := strings.TrimSpace(
		terminalinput.ReadInput("Enter filepath\n\r-> "))
	column, err := strconv.Atoi(
		strings.TrimSpace(
			terminalinput.ReadInput("Enter column of csv containing the email addresses\n\r-> ")))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	list, err := parser.Csv(filePath, column)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if err := database.MailingListAdd(name, list); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

func cmdAdd_Message() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter a name\n\r-> "))
	subject := strings.TrimSpace(
		terminalinput.ReadInput("Enter a subject for this message\n\r-> "))
	filePath := strings.TrimSpace(
		terminalinput.ReadInput("Enter the message. Either .html or .txt\n\r-> "))
	message, msg_type, err := parser.Email(filePath)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if err := database.MessageAdd(name, subject, msg_type, message); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}
