package database

import (
  "errors"
)

func RecipientsAdd(name string, list []string) error {
  Db, err := openDatabase()
  if err != nil {
    return err
  }

  if _, exists := Db.Recipients[name]; exists != false {
    return errors.New("Recipients list already present.")
  }
  Db.Recipients[name] = list

  err = writeDatabase(Db)
  if err != nil {
   return err
  }
  return nil
}

func RecipientsDelete(name string) error {
  Db, err := openDatabase()
  if err != nil {
    return err
  }

  delete(Db.Recipients, name)

  err = writeDatabase(Db)
  if err != nil {
    return err
  }
  return nil
}
