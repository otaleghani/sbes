package repl

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

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