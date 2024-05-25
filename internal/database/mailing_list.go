package database

import (
	"errors"
  "fmt"
)

func MailingListAdd(name string, list []string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	if _, exists := Db.Mailing_lists[name]; exists != false {
		return errors.New("Mailing list already present.")
	}
	Db.Mailing_lists[name] = list

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MailingListDelete(name string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	delete(Db.Mailing_lists, name)

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MailingListsList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}
  if len(Db.Mailing_lists) != 0 {
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  for index, val := range Db.Mailing_lists {
    fmt.Printf("%v | 5 items: %v\n", index, val[0] + val[1] + val[2])
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  return nil
}

func MailingListGet(name string) (string, []string, error) {
  // name, subject, msg_type, body
	Db, err := openDatabase()
	if err != nil {
	  return "", []string{}, errors.New("Record does not exist")
	}

  val, exists := Db.Mailing_lists[name]
  if exists != true {
	  return "", []string{}, errors.New("Record does not exist")
  }

	return name, val, nil
}
