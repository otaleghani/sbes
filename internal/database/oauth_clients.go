package database

import (
  "errors"
)

func OauthClientAdd(name, client_id, client_secret string) error {
  Db, err := openDatabase()
  if err != nil {
    return err
  }

  if _, exists := Db.OAuth_clients[name]; exists != false {
    return errors.New("OAuth client already present. Delete it first.")
  }
  Db.OAuth_clients[name] = OAuthClient{
    Client_Id: client_id,
    Client_Secret: client_secret,
    Name: name,
  }

  err = writeDatabase(Db)
  if err != nil {
   return err
  }
  return nil
}

func OauthClientDelete(name string) error {
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
