// Functions to manupulate the accounts inside of the database

package database

import (
	"errors"
	"fmt"
)

func AccountAdd(user, pass, host string, port int) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Checks if user already exist
	if _, exists := Db.Accounts[user]; exists {
		return errors.New("username already present")
	}

	// Adds account
	Db.Accounts[user] = Account{
		Username: user, Password: pass, SmtpHost: host, SmtpPort: port,
		RefreshToken: "", AccessToken: ""}

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func AccountDelete(user string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Deletes account
	delete(Db.Accounts, user)

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func AccountsList() error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Prints accounts
	fmt.Printf(
		"%-30v %-30v %6v %6v\n",
		"USERNAME", "HOST", "PORT", "OAUTH")
	for name, val := range Db.Accounts {
		oauth := "NO"
		if val.RefreshToken != "" && val.AccessToken != "" {
			oauth = "YES"
		}
		truncIndex := name
		truncHost := val.SmtpHost
		if len(truncIndex) > 29 {
			truncIndex = truncIndex[:29]
		}
		if len(truncHost) > 29 {
			truncHost = truncHost[:29]
		}
		fmt.Printf(
			"%-30v %-30v %6v %6v\n",
			truncIndex, truncHost, val.SmtpPort, oauth)
	}
	return nil
}

func AccountGet(user string) (string, string, string, string, string, int, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", "", "", "", "", 0, nil
	}

	// Checks if record exists
	val, exists := Db.Accounts[user]
	if !exists {
		return "", "", "", "", "", 0, errors.New("record does not exist")
	}

	// Returns values of account
	return val.Username, val.Password, val.SmtpHost, val.RefreshToken, val.AccessToken, val.SmtpPort, nil
}

func AccountTokenAdd(user, refreshToken, accessToken string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Checks if record exists
	_, exists := Db.Accounts[user]
	if !exists {
		return errors.New("user not found")
	}

	// Adds access and refresh token
	account := Db.Accounts[user]
	account.AccessToken = accessToken
	account.RefreshToken = refreshToken
	Db.Accounts[user] = account

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}
