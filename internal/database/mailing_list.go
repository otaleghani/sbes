package database

import (
	"errors"
	"fmt"
	"strings"
)

func MailingListAdd(name string, list []string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	if _, exists := Db.MailingLists[name]; exists != false {
		return errors.New("Mailing list already present.")
	}
	Db.MailingLists[name] = MailingList{List: list}

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

	delete(Db.MailingLists, name)

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
	fmt.Printf(
		"%-20v %-55v\n",
		"NAME", "EXAMPLE")

	for index, val := range Db.MailingLists {
		truncIndex := index
		if len(truncIndex) > 19 {
			truncIndex = truncIndex[:19]
		}

		items := len(val.List)
		if len(val.List) > 5 {
			items = 5
		}
		examplesArray := []string{}

		for i := 0; i < items; i++ {
			examplesArray = append(examplesArray, val.List[i])
		}
		exampleString := strings.Join(examplesArray, ", ")

		truncExamples := exampleString
		if len(exampleString) > 54 {
			truncExamples = exampleString[:54]
		}
		fmt.Printf("%-20v %-55v\n", index, truncExamples)
	}
	return nil
}

func MailingListGet(name string) (string, []string, error) {
	// name, subject, msg_type, body
	Db, err := openDatabase()
	if err != nil {
		return "", []string{}, errors.New("Record does not exist")
	}

	val, exists := Db.MailingLists[name]
	if exists != true {
		return "", []string{}, errors.New("Record does not exist")
	}

	return name, val.List, nil
}
