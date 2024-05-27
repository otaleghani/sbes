package repl

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/sender"
	"github.com/otaleghani/sbes/internal/terminalinput"
)

func cmdSend_password() {
	user := strings.TrimSpace(
		terminalinput.ReadInput("Choose the account\n\r-> "))
	user, pass, host, port, _, _, err := database.AccountGet(user)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println(user, pass, host, port)

	message := strings.TrimSpace(
		terminalinput.ReadInput("Choose the message\n\r-> "))
	messageName, subject, msg_type, body, err := database.MessageGet(message)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println(messageName, subject, msg_type, body)

	mailing_list := strings.TrimSpace(
		terminalinput.ReadInput("Choose the mailing list\n\r-> "))
	mailingListName, list, err := database.MailingListGet(mailing_list)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println(mailingListName, list)

	// get account, message and mailing list, then call SendWithPassword()
}

func cmdSend_oauth() {
	account := strings.TrimSpace(
		terminalinput.ReadInput("Choose the account\n\r-> "))

	user, pass, host, refresh, access, port, err := database.AccountGet(account)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if refresh == "" {
		fmt.Println("ERROR: Refresh token is empty")
		return
	}
	if access == "" {
		fmt.Println("You will need to add a new access token to this account.")
		return
	}
	fmt.Println(user, pass, host, port, refresh, access)

	message := strings.TrimSpace(
		terminalinput.ReadInput("Choose the message\n\r-> "))
	messageName, subject, msg_type, body, err := database.MessageGet(message)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(messageName, subject, msg_type, body)

	mailing_list := strings.TrimSpace(
		terminalinput.ReadInput("Choose the mailing list\n\r-> "))
	mailingListName, list, err := database.MailingListGet(mailing_list)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(mailingListName, list)

	// get account, message and mailing list, then call SendWithPassword()
	email := sender.Email{
		// Authentication
		SmtpHost: host,
		SmtpPort: port,
		Username: account,
		Password: pass,
		Oauth:    access,
		// Message
		From:         account,
		Mailing_List: list,
		Subject:      subject,
		Body:         body,
		Msg_Type:     msg_type,
	}
	sender.SendEmailOAuth(email)
}

func cmdSend() {

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	recipients := sendCmd.String("r", "Recipients list", "Path to the recipients .csv file.")
	auth := sendCmd.String("a", "Default Body", "Credentials added by using \"sbes auth\"")
	message := sendCmd.String("m", "Default Body", "Message to be send. Either .txt for plaintext or .html")
	sendCmd.Parse(os.Args[2:])

	err := handleCmdSend(*recipients, *auth, *message)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func handleCmdSend(recipients, auth, message string) error {
	// Db, err := openDatabase()
	// account, exists := Db.Accounts[auth]
	// if exists != true {
	//   return errors.New("ERROR: No account found")
	// }

	// Parses recipients list
	csvFile, err := os.Open(recipients)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	var data []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New("ERROR: EOF")
		}

		// Check if the specified column exists in the record
		if 0 < len(record) {
			data = append(data, record[0])
		} else {
			return errors.New("ERROR: column")
		}
	}

	data = data[1:]

	// second username is from field
	// err = SendEmails(
	//   account.SmtpHost,
	//   account.SmtpPort,
	//   "username",
	//   account.Password,
	//   account.Username,
	//   data,
	//   "testing",
	//   message)
	// if err != nil {
	//   return err
	// }

	fmt.Println("Email sent successfully!")
	return nil
}
