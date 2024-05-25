package repl

import (
	"os"
	"strings"
)

func Start() {
	if len(os.Args) == 1 {
		cmdHelp()
		return
	}
	cmd := strings.ToLower(os.Args[1])
	switch cmd {

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
			cmdAdd_OAuthToken()
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
		case "oauth-token":
      cmdList_OAuthTokens()
		case "mailing-list":
      cmdList_MailingLists()
		case "message":
      cmdList_Messages()
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
		case "oauth-token":
      cmdDelete_OAuthToken()
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
		cmdHelp()

	default:
		cmdHelp()
	}
}
