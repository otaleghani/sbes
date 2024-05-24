package database

import (
	"errors"
)

func OAuthTokenAdd(name, token string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	if _, exists := Db.OAuth_tokens[name]; exists != false {
		return errors.New("OAuth token already present. Delete it first.")
	}
	Db.OAuth_tokens[name] = OAuthToken{
		Username: name,
		Token:    token,
	}

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func OAuthTokenDelete(name string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	delete(Db.OAuth_tokens, name)

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}
