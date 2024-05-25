package database

import (
	"errors"
  "fmt"
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

func OAuthTokensList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}
  if len(Db.OAuth_tokens) != 0 {
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  for index, val := range Db.OAuth_tokens {
    fmt.Printf("%v | Username: %v | Token: %v\n", index, val.Token, val.Token)
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  return nil
}

func OAuthTokenGet(name string) (string, string, error) {
  // name, subject, msg_type, body
	Db, err := openDatabase()
	if err != nil {
	  return "", "", errors.New("Record does not exist")
	}

  val, exists := Db.OAuth_tokens[name]
  if exists != true {
	  return "", "", errors.New("Record does not exist")
  }

	return val.Username, val.Token, nil
}
