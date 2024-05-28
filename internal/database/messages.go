// Functions to manupulate the messages inside the database

package database

import (
	"errors"
	"fmt"
)

func MessageAdd(name, subject, msg_type, body string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Checks if the message is already present
	if _, exists := Db.Messages[name]; exists {
		return errors.New("message already present")
	}

	// Adds the message
	Db.Messages[name] = Message{
		Subject: subject,
		Body:    body,
		MsgType: msg_type,
	}

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MessageDelete(name string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Deletes the message
	delete(Db.Messages, name)

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MessagesList() error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Prints data
	fmt.Printf(
		"%-20v %-47v %6v\n",
		"NAME", "SUBJECT", "TYPE")

	// Cycles through every item in the db
	for name, val := range Db.Messages {
		// Truncates the name if > 19
		truncName := name
		if len(truncName) > 19 {
			truncName = truncName[:19]
		}

		// Truncates the subject if > 19
		truncSubject := val.Subject
		if len(truncSubject) > 45 {
			truncSubject = truncSubject[:45]
		}

		// Prints the row
		fmt.Printf(
			"%-20v %-47v %6v\n",
			truncName, truncSubject, val.MsgType)
	}
	return nil
}

func MessageGet(name string) (string, string, string, string, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", "", "", "", nil
	}

	// Checks if the message is present
	val, exists := Db.Messages[name]
	if !exists {
		return "", "", "", "", errors.New("record does not exist")
	}

	// Returns it
	return name, val.Subject, val.MsgType, val.Body, nil
}
