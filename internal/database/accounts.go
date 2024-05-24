package database

import (
  "errors"
)

func AccountAdd(username, password, smtpHost string, smtpPort int) error {
  Db, err := openDatabase()
  if err != nil {
    return err
  }

  // Checks if account is present
  if _, exists := Db.Accounts[username]; exists != false {
    return errors.New("Username already present. If you want to update it use auth-update")
  }
  Db.Accounts[username] = Account{
    Username: username,
    Password: password,
    SmtpHost: smtpHost,
    SmtpPort: smtpPort,
  }

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
