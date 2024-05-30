// Functions to manupulate the mailing lists inside the database

package database

import (
	"errors"
	"fmt"
	"strings"
)

func MailingListAdd(name string, list []string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Checks if the email list is alreay present
	if _, exists := Db.MailingLists[name]; exists {
		return errors.New("mailing list already present")
	}
	Db.MailingLists[name] = MailingList{
    Name: name,
    List: list,
  }

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MailingListDelete(name string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Deletes item
	delete(Db.MailingLists, name)

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func MailingListsList() error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Prints data
	fmt.Printf(
		"%-20v %-55v\n",
		"NAME", "EXAMPLE")

	// Cycles through every item
	for name, val := range Db.MailingLists {
		// Truncates name if > 19 chars
		truncName := name
		if len(truncName) > 19 {
			truncName = truncName[:19]
		}

		// Takes 5 or less items
		items := len(val.List)
		if len(val.List) > 5 {
			items = 5
		}
		examplesArray := []string{}

		// Cicles throght the List of email and appends them for the lenght of items
		for i := 0; i < items; i++ {
			examplesArray = append(examplesArray, val.List[i])
		}
		// Joins the emails into a string
		exampleString := strings.Join(examplesArray, ", ")

		// Truncates the emails if > 54 chars
		truncExamples := exampleString
		if len(exampleString) > 54 {
			truncExamples = exampleString[:54]
		}

		// Prints the row
		fmt.Printf("%-20v %-55v\n", truncName, truncExamples)
	}
	return nil
}

func MailingListGet(name string) (string, []string, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", []string{}, errors.New("record does not exist")
	}

	// Checks if the item exists
	val, exists := Db.MailingLists[name]
	if !exists {
		return "", []string{}, errors.New("record does not exist")
	}

	// Returns name and email list
	return name, val.List, nil
}

func MailingListGetObject(name string) (MailingList, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return MailingList{}, nil
	}

	// Checks if record exists
	val, exists := Db.MailingLists[name]
	if !exists {
		return MailingList{}, errors.New("record does not exist")
	}

	// Returns values of account
	return val, nil
}
