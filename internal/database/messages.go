package database

import (
	"errors"
	"fmt"
)

func MessageAdd(name, subject, msg_type, body string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	if _, exists := Db.Messages[name]; exists != false {
		return errors.New("Message already present.")
	}

	Db.Messages[name] = Message{
		Subject: subject,
		Body:    body,
		MsgType: msg_type,
	}

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MessageDelete(name string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	delete(Db.Messages, name)

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MessagesList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	fmt.Printf(
		"%-20v %-47v %6v\n",
		"NAME", "SUBJECT", "TYPE")

	for name, val := range Db.Messages {
		truncName := name
		if len(truncName) > 19 {
			truncName = truncName[:19]
		}

		truncSubject := val.Subject
		if len(truncSubject) > 45 {
			truncSubject = truncSubject[:45]
		}

		fmt.Printf(
			"%-20v %-47v %6v\n",
			truncName, truncSubject, val.MsgType)
	}
	return nil
}

func MessageGet(name string) (string, string, string, string, error) {
	// name, subject, msg_type, body
	Db, err := openDatabase()
	if err != nil {
		return "", "", "", "", nil
	}

	val, exists := Db.Messages[name]
	if exists != true {
		return "", "", "", "", errors.New("Record does not exist")
	}

	return name, val.Subject, val.MsgType, val.Body, nil
}
