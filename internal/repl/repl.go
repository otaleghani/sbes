package repl

import (
	"fmt"
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
			fmt.Println("list account")
		case "oauth-client":
			fmt.Println("list oauth client")
		case "oauth-token":
			fmt.Println("list oauth token")
		case "mailing-list":
			fmt.Println("list mailing-list")
		case "message":
			fmt.Println("list message")
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
			fmt.Println("delete account")
		case "oauth-client":
			fmt.Println("delete oauth client")
		case "oauth-token":
			fmt.Println("delete oauth token")
		case "mailing-list":
			fmt.Println("delete mailing-list")
		case "message":
			fmt.Println("delete message")
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
			fmt.Println("password based auth")
		case "oauth":
			fmt.Println("oauth authentication")
		}

	case "help":
		cmdHelp()

	default:
		cmdHelp()
	}
}

func cmdHelp() {
	fmt.Println("Help ")
}
