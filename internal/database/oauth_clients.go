// Functions to manupulate the oauth clients inside the database

package database

import (
	"errors"
	"fmt"
)

func OAuthClientAdd(name, client_id, client_secret string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Checks if the item is present
	if _, exists := Db.OAuthClients[name]; exists {
		return errors.New("oauth client already present, delete it first")
	}

	// Adds the item
	Db.OAuthClients[name] = OAuthClient{
		ClientId:     client_id,
		ClientSecret: client_secret,
	}

	// Closes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func OAuthClientDelete(name string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Deletes the item
	delete(Db.OAuthClients, name)

	// Closes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func OAuthClientsList() error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Prints header
	fmt.Printf(
		"%-20v %-26v %-28v\n",
		"NAME", "ID", "SECRET")

	// Cycles through every item
	for name, val := range Db.OAuthClients {
		// Truncates name if > 19
		truncName := name
		if len(truncName) > 19 {
			truncName = truncName[:19]
		}

		// Truncates id if > 25
		truncId := val.ClientId
		if len(truncId) > 25 {
			truncId = truncId[:25]
		}

		// Truncates secret if > 27
		truncSecret := val.ClientSecret
		if len(truncSecret) > 27 {
			truncSecret = truncSecret[:27]
		}

		// Prints row
		fmt.Printf(
			"%-20v %-26v %-28v\n",
			truncName, truncId, truncSecret)
	}
	return nil
}

func OauthClientGet(name string) (string, string, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", "", errors.New("record does not exist")
	}

	// Checks if the record exists
	val, exists := Db.OAuthClients[name]
	if !exists {
		return "", "", errors.New("record does not exist")
	}

	// Returns id and secret
	return val.ClientId, val.ClientSecret, nil
}
