package repl

import (
	"fmt"
	"os"
	"strings"
)

// Starts the repl
func Start() {
	// Checks if there is more than one arg
	if len(os.Args) == 1 {
		cmdHelp()
		return
	}

	// Formats the cmd
	cmd := strings.ToLower(os.Args[1])

	// Searches for inserted cmd, kinda self explantory
	switch cmd {
	case "refresh":
		cmdRefreshAccessToken()

	case "add":
		if len(os.Args) == 2 {
			cmdHelp()
			return
		}
		item := strings.ToLower(os.Args[2])
		switch item {
		case "account":
			cmdAdd_Account()
		case "oauth-client":
			cmdAdd_OAuthClient()
		case "oauth-token":
			cmdAdd_OAuthRefreshToken()
		case "mailing-list":
			cmdAdd_MailingList()
		case "message":
			cmdAdd_Message()
		default:
			cmdHelp()
		}

	case "list":
		if len(os.Args) == 2 {
			cmdHelp()
			return
		}
		item := strings.ToLower(os.Args[2])
		switch item {
		case "account":
			cmdList_Accounts()
		case "oauth-client":
			cmdList_OAuthClients()
		case "mailing-list":
			cmdList_MailingLists()
		case "message":
			cmdList_Messages()
		case "campaign":
			cmdList_Campaigns()
		default:
			cmdHelp()
		}

	case "delete":
		if len(os.Args) == 2 {
			cmdHelp()
			return
		}
		item := strings.ToLower(os.Args[2])
		switch item {
		case "account":
			cmdDelete_Account()
		case "oauth-client":
			cmdDelete_OAuthClient()
		case "mailing-list":
			cmdDelete_MailingList()
		case "message":
			cmdDelete_Message()
		default:
			cmdHelp()
		}

	case "send":
		if len(os.Args) == 2 {
			cmdHelp()
			return
		}
		// Login method...
		auth := strings.ToLower(os.Args[2])
		switch auth {
		case "password":
			cmdSend_password()
		case "oauth":
			cmdSend_oauth()
		default:
			cmdHelp()
		}

	case "help":
		if len(os.Args) == 2 {
			cmdHelp()
			return
		}
		help := strings.ToLower(os.Args[2])
		switch help {
		case "add":
			cmdHelpAdd()
		case "delete":
			cmdHelpDelete()
		case "list":
			cmdHelpList()
		case "send":
			cmdHelpSend()
		default:
			cmdHelp()
		}

	default:
		cmdHelp()
	}
}

func divider() {
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
}
