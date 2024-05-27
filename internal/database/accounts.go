package database

import (
	"errors"
	"fmt"
)

func AccountAdd(user, pass, host string, port int) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}
	if _, exists := Db.Accounts[user]; exists != false {
		return errors.New("Username already present.")
	}
	Db.Accounts[user] = Account{
		Username: user, Password: pass, SmtpHost: host, SmtpPort: port,
		RefreshToken: "", AccessToken: ""}

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func AccountDelete(user string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	delete(Db.Accounts, user)

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func AccountsList() error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

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
	Db, err := openDatabase()
	if err != nil {
		return "", "", "", "", "", 0, nil
	}

	val, exists := Db.Accounts[user]
	if exists != true {
		return "", "", "", "", "", 0, errors.New("Record does not exist")
	}

	return val.Username, val.Password, val.SmtpHost, val.RefreshToken, val.AccessToken, val.SmtpPort, nil
}

func AccountTokenAdd(user, refreshToken, accessToken string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	_, exists := Db.Accounts[user]
	if exists != true {
		return errors.New("User not found.")
	}

	account := Db.Accounts[user]
	account.AccessToken = accessToken
	account.RefreshToken = refreshToken
	Db.Accounts[user] = account

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}
