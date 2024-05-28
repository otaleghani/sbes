package repl

import (
	"fmt"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/sender"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdSend_password() {
	cmdList_Accounts()
	user := strings.TrimSpace(
		terminalinput.ReadInput("Choose the account\n\r-> "))
	user, pass, host, _, _, port, err := database.AccountGet(user)
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println()

	cmdList_Messages()
	message := strings.TrimSpace(
		terminalinput.ReadInput("Choose the message\n\r-> "))
	_, subject, msg_type, body, err := database.MessageGet(message)
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println()

	cmdList_MailingLists()
	mailing_list := strings.TrimSpace(
		terminalinput.ReadInput("Choose the mailing list\n\r-> "))
	_, list, err := database.MailingListGet(mailing_list)
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println()
	divider()

	err = sender.SendEmails(sender.Email{
		SmtpHost:    host,
		SmtpPort:    port,
		Username:    user,
		Password:    pass,
		Oauth:       "",
		From:        user,
		MailingList: list,
		Subject:     subject,
		Body:        body,
		MsgType:     msg_type,
	})
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
}

func cmdSend_oauth() {
	cmdList_Accounts()
	account := strings.TrimSpace(
		terminalinput.ReadInput("Choose the account\n\r-> "))
	_, pass, host, refresh, access, port, err := database.AccountGet(account)
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
	if refresh == "" {
		divider()
		fmt.Println("ERROR: Refresh token is empty")
		return
	}
	if access == "" {
		divider()
		fmt.Println("ERROR: Access token not found, update it with sbes add access-token")
		return
	}
	fmt.Println()

	cmdList_Messages()
	message := strings.TrimSpace(
		terminalinput.ReadInput("Choose the message\n\r-> "))
	_, subject, msg_type, body, err := database.MessageGet(message)
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println()

	cmdList_MailingLists()
	mailing_list := strings.TrimSpace(
		terminalinput.ReadInput("Choose the mailing list\n\r-> "))
	_, list, err := database.MailingListGet(mailing_list)
	if err != nil {
		divider()
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println()
	divider()

	// get account, message and mailing list, then call SendWithPassword()
	email := sender.Email{
		// Authentication
		SmtpHost: host,
		SmtpPort: port,
		Username: account,
		Password: pass,
		Oauth:    access,
		// Message
		From:        account,
		MailingList: list,
		Subject:     subject,
		Body:        body,
		MsgType:     msg_type,
	}
	sender.SendEmailOAuth(email)
}
