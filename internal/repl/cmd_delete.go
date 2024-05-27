package repl

import (
	"fmt"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdDelete_Account() {
	user := strings.TrimSpace(
		terminalinput.ReadInput("Enter username that you want to delete\n\r-> "))

	if err := database.AccountDelete(user); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

func cmdDelete_OAuthClient() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name of OAuthClient that you want to delete\n\r-> "))
	if err := database.OAuthClientDelete(name); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

// func cmdDelete_OAuthToken() {
// 	name := strings.TrimSpace(
// 		terminalinput.ReadInput("Enter name of OAuthToken that you want to delete\n\r-> "))
// 	if err := database.OAuthTokenDelete(name); err != nil {
// 		fmt.Println("ERROR: ", err)
// 		return
// 	}
// }

func cmdDelete_MailingList() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name of mailing list that you want to delete\n\r-> "))
	if err := database.MailingListDelete(name); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}

func cmdDelete_Message() {
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name of message that you want to delete\n\r-> "))
	if err := database.MessageDelete(name); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}
