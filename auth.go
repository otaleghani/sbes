package main

import (
  "encoding/json"
  "os"
  // "path/filepath"
  "fmt"
  "errors"
)

type Database struct {
  Accounts map[string]string `json:accounts`
}

func AddAuth(username, password string) error {
  homePath, err := os.UserHomeDir()
  if err != nil {
    return err
  }
  path := homePath + "/.cache/sbes/"
  fileName := path + "accounts.json"

  if _, err := os.Stat(fileName); os.IsNotExist(err) {
    err = os.MkdirAll(path, 0755)
    if err != nil {
      return err
    }
    err = os.WriteFile(fileName, []byte("{}"), 0666)
    if err != nil {
      return err
    }
  }

  // Check if account is already present
  data, err := os.ReadFile(fileName)
  if err != nil {
    return err
  }

  Database := Database{Accounts: make(map[string]string)}
  if err = json.Unmarshal(data, &Database); err != nil {
    fmt.Println(data)
    return err
  }

  if Database.Accounts[username] == password {
    return errors.New("Username already present")
  }

  Database.Accounts[username] = password
  
  encodedData, err := json.Marshal(Database)

  os.WriteFile(fileName, encodedData, 0666)

  return nil
}
