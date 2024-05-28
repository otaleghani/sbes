package repl

import (
	"fmt"
	"github.com/otaleghani/sbes/internal/database"
)

func cmdList_Accounts() {
	// Calls function to list all accounts
	divider()
	if err := database.AccountsList(); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	divider()
}

func cmdList_OAuthClients() {
	// Calls function to list all OAuth Clients
	divider()
	if err := database.OAuthClientsList(); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	divider()
}

func cmdList_MailingLists() {
	// Calls function to list all mailing lists
	divider()
	if err := database.MailingListsList(); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	divider()
}

func cmdList_Messages() {
	// Calls function to list all messages
	divider()
	if err := database.MessagesList(); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	divider()
}
