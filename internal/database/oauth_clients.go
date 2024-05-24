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

	if _, exists := Db.OAuth_clients[name]; exists != false {
		return errors.New("OAuth client already present. Delete it first.")
	}
	Db.OAuth_clients[name] = OAuthClient{
		Client_Id:     client_id,
		Client_Secret: client_secret,
		Name:          name,
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

	delete(Db.OAuth_clients, name)

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

	if _, exists := Db.OAuth_clients[name]; exists != true {
		return "", "", errors.New("OAuth does not exist.")
	}

	return Db.OAuth_clients[name].Client_Id, Db.OAuth_clients[name].Client_Secret, nil
}

func OAuthClientsList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}
  if len(Db.OAuth_clients) != 0 {
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  for index, val := range Db.OAuth_clients {
    fmt.Printf("%v | Id: %v | Secret: %v\n", index, val.Client_Id, val.Client_Secret)
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  return nil
}
