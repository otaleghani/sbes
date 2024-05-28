package repl

import (
	"fmt"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdDelete_Account() {
  // Propts the user to select an account
  cmdList_Accounts()
	user := strings.TrimSpace(
		terminalinput.ReadInput("Enter username that you want to delete\n\r-> "))
  divider()

  // Tries to delete said account
	if err := database.AccountDelete(user); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Account deleted from database.")
}

func cmdDelete_OAuthClient() {
  // Propts the user to select an OAuth Client
  cmdList_OAuthClients()
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name of OAuthClient that you want to delete\n\r-> "))
  divider()

  // Tries to delete said Oauth Client
	if err := database.OAuthClientDelete(name); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: OAuth2 client deleted from database.")
}

func cmdDelete_MailingList() {
  // Propts the user to select an mailing list
  cmdList_MailingLists()
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name of mailing list that you want to delete\n\r-> "))
  divider()

  // Tries to delete said account
	if err := database.MailingListDelete(name); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Mailing list deleted from database.")
}

func cmdDelete_Message() {
  // Propts the user to select a message
  cmdList_Messages()
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name of message that you want to delete\n\r-> "))
  divider()

  // Tries to delete said message
	if err := database.MessageDelete(name); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Message deleted from database.")
}
