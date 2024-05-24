package database

import (
	"errors"
)

func MessageAdd(name, subject, msg_type string, body string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	if _, exists := Db.Messages[name]; exists != false {
		return errors.New("Message already present.")
	}


	Db.Messages[name] = Message{
		Subject:  subject,
		Body:     body,
		Msg_Type: msg_type,
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
