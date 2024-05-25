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
	// Checks if account is present
	if _, exists := Db.Accounts[user]; exists != false {
		return errors.New("Username already present. If you want to update it use auth-update")
	}
	Db.Accounts[user] = Account{
    Username:user, Password:pass, SmtpHost:host, SmtpPort:port}

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func AccountDelete(username string) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	delete(Db.Accounts, username)

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func AccountUpdate(account Account) error {
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	Db.Accounts[account.Username] = account

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
  if len(Db.Accounts) != 0 {
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  for index, val := range Db.Accounts {
    fmt.Printf("%v | Host: %v | Port: %v\n", index, val.SmtpHost, val.SmtpPort)
    fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  }
  return nil
}

func AccountGet(user string) (string, string, string, int, error) {
	Db, err := openDatabase()
	if err != nil {
	  return "","","",0,nil
	}

  val, exists := Db.Accounts[user]
  if exists != true {
	  return "", "", "", 0, errors.New("Record does not exist")
  }

	return val.Username, val.Password, val.SmtpHost, val.SmtpPort, nil
}
