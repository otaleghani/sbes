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

func MessagesList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}
  if len(Db.Messages) != 0 {
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  for index, val := range Db.Messages {
   fmt.Printf("%v | Type: %v | Body: %v\n", index, val.Msg_Type, val.Body)
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
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

	return name, val.Subject, val.Msg_Type, val.Body, nil
}
