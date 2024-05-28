// Helper REPL functions add items inside the database

package repl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/oauth2"
	"github.com/otaleghani/sbes/internal/parser"
	"github.com/otaleghani/sbes/internal/sender"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdAdd_Account() {
  // Gets name of user, checks if already present
  divider()
	user := strings.TrimSpace(
		terminalinput.ReadInput("Enter the username\n\r-> "))
	username, _, _, _, _, _, err := database.AccountGet(user)
	if username != "" {
		fmt.Println("ERROR: User already present")
		return
	}
  divider()

  // Gets password
	pass := strings.TrimSpace(
		terminalinput.ReadInput("Enter the password\n\r-> "))
  divider()

  // Gets hostname
	host := strings.TrimSpace(
		terminalinput.ReadInput("Enter the hostname\n\r-> "))
  divider()

  // Gets port, converting it in a int
	port, err := strconv.Atoi(
		strings.TrimSpace(
			terminalinput.ReadInput("Enter the port\n\r-> ")))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
  divider()

	// Tests connection with server
	fmt.Println("LOG: Testing connection...")
	err = sender.TestSMTPConnection(host, port, user, pass)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Connection successful")

  // Saves data inside of the database
	fmt.Println("LOG: Saving data to local database")
	if err := database.AccountAdd(user, pass, host, port); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Account saved")
}

func cmdAdd_OAuthClient() {
  // Gets name of OAuth Client to add, checks if already present
  divider()
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter a name for this OAuth Client\n\r-> "))
	id, _, err := database.OauthClientGet(user)
	if id != "" {
		fmt.Println("ERROR: An OAuth Client is already present with this name")
		return
	}
  divider()
  
  // Gets id
	id := strings.TrimSpace(
		terminalinput.ReadInput("Enter OAuth Client Id\n\r-> "))
  divider()

  // Gets secret
	secret := strings.TrimSpace(
		terminalinput.ReadInput("Enter OAuth Client Secret\n\r-> "))
  divider()

  // Adds OAuth to database
	if err := database.OAuthClientAdd(name, id, secret); err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("SUCCESS: OAuth Client saved")
}


func cmdAdd_MailingList() {
  // Gets name, checks if already present
  divider()
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter name for mailing list\n\r-> "))
	id, _, err := database.MailingListGet(name)
	if id != "" {
		fmt.Println("ERROR: Email list already present with this name.")
    return
  }
  divider()

  // Gets file path
	filePath := strings.TrimSpace(
		terminalinput.ReadInput("Enter filepath\n\r-> "))
  divider()

  // Gets column to scan, convers it to int
	column, err := strconv.Atoi(
		strings.TrimSpace(
			terminalinput.ReadInput("Enter column of csv containing the email addresses\n\r-> ")))
  divider()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

  // Calls parser to parse the Csv
	list, err := parser.Csv(filePath, column)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

  // Writes data the database
	if err := database.MailingListAdd(name, list); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Mailing list added to database.")
}

func cmdAdd_Message() {
  // Adds name of message, checks if already exists
  divider()
	name := strings.TrimSpace(
		terminalinput.ReadInput("Enter a name\n\r-> "))
	id, _, _, _, err := database.MessageGet(name)
	if id != "" {
		fmt.Println("ERROR: Message with this name is already present.")
    return
  }
  divider()

  // Gets subject of email
	subject := strings.TrimSpace(
		terminalinput.ReadInput("Enter a subject for this message\n\r-> "))
  divider()

  // Gets filepath of file
	filePath := strings.TrimSpace(
		terminalinput.ReadInput("Enter the message. Either .html or .txt\n\r-> "))
  divider()

  // Calls parser for that filepath
	message, msg_type, err := parser.Email(filePath)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

  // Adds message to database
	if err := database.MessageAdd(name, subject, msg_type, message); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
  fmt.Println("SUCCESS: Message added to database.")
}


func cmdAdd_OAuthRefreshToken() {
  // Propts the user to select a user
  cmdList_Accounts()
	user := strings.TrimSpace(
		terminalinput.ReadInput("Account username\n\r-> "))
	username, _, _, _, _, _, err := database.AccountGet(user)
	if err != nil {
    divider()
		fmt.Println("ERROR: ", err)
		return
	}
  fmt.Println()

  // Propts the user to select a OAuth Client
  cmdList_OAuthClients()
	nameClient := strings.TrimSpace(
		terminalinput.ReadInput("Client to generate token\n\r-> "))
	id, secret, err := database.OauthClientGet(nameClient)
	if err != nil {
    divider()
		fmt.Println("ERROR: ", err)
		return
	}
  fmt.Println()

  // Gets refresh and access token
	refreshToken, accessToken := oauth2.GetOauth2(id, secret)
	if refreshToken == "" {
    divider()
		fmt.Println("ERROR: Token not generated. Check you client credentials.")
		return
	}

  // Writes changes to selected user
	if err := database.AccountTokenAdd(username, refreshToken, accessToken); err != nil {
    divider()
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: OAuth Client saved")
}

// Refreshes the access token
func cmdRefreshAccessToken() {
  // Propts the user to select a user
  cmdList_Accounts()
	account := strings.TrimSpace(
		terminalinput.ReadInput("Choose the account\n\r-> "))
	user, _, _, refresh, _, _, err := database.AccountGet(account)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
  fmt.Println()

  // Prompts the user to select a OAuth Client
  cmdList_OAuthClients()
  client := strings.TrimSpace(
		terminalinput.ReadInput("What client do you want to use?\n\r-> "))
	id, secret, err := database.OauthClientGet(client)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
  divider()

  // Gets new Access token
  access, err := oauth2.GetAccessToken(id, secret, refresh)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

  // Saves it to database
  err = database.AccountTokenAdd(user, refresh, access)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("SUCCESS: Access token refreshed correctly.")
}
