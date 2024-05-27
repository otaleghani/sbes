package database

import (
	"errors"
	"fmt"
)

func OAuthClientAdd(name, client_id, client_secret string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	if _, exists := Db.OAuthClients[name]; exists != false {
		return errors.New("OAuth client already present. Delete it first.")
	}
	Db.OAuthClients[name] = OAuthClient{
		ClientId:     client_id,
		ClientSecret: client_secret,
	}

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func OAuthClientDelete(name string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	delete(Db.OAuthClients, name)

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func OAuthClientFind(name string) (string, string, error) {
	Db, err := openDatabase()
	if err != nil {
		return "", "", err
	}

	if _, exists := Db.OAuthClients[name]; exists != true {
		return "", "", errors.New("OAuth does not exist.")
	}

	return Db.OAuthClients[name].ClientId, Db.OAuthClients[name].ClientSecret, nil
}

func OAuthClientsList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	fmt.Printf(
		"%-20v %-26v %-28v\n",
		"NAME", "ID", "SECRET")

	for name, val := range Db.OAuthClients {
		truncName := name
		if len(truncName) > 19 {
			truncName = truncName[:19]
		}
		truncId := val.ClientId
		if len(truncId) > 25 {
			truncId = truncId[:25]
		}
		truncSecret := val.ClientSecret
		if len(truncSecret) > 27 {
			truncSecret = truncSecret[:27]
		}

		fmt.Printf(
			"%-20v %-26v %-28v\n",
			truncName, truncId, truncSecret)
	}
	return nil
}

func OauthClientGet(name string) (string, string, error) {
	// name, subject, msg_type, body
	Db, err := openDatabase()
	if err != nil {
		return "", "", errors.New("Record does not exist")
	}

	val, exists := Db.OAuthClients[name]
	if exists != true {
		return "", "", errors.New("Record does not exist")
	}

	return val.ClientId, val.ClientSecret, nil
}
