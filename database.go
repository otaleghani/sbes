package main

import (
  "os"
  "encoding/json"
  "errors"
)

// To do
// Get account

func getPaths() (string, string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	path := homePath + "/.cache/sbes/"
	fileName := path + "accounts.json"
  
  return path, fileName, nil
}

func openDatabase() (Database, error) {
  path, fileName, err := getPaths()
  if err != nil {
    return Database{}, err
  }

  // Tests if the file is present
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return Database{}, err
		}
		err = os.WriteFile(fileName, []byte("{}"), 0666)
		if err != nil {
			return Database{}, err
		}
	}

  // Parses data from file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return Database{}, err
	}

  // Marshals the json
  Db := Database{Accounts: make(map[string]Config)}
  if err = json.Unmarshal(data, &Db); err != nil {
    return Database{}, err
  }
  return Db, nil
}

func writeDatabase(db Database) error {
  _, fileName, err := getPaths()
  if err != nil {
    return err
  }

  encodedData, err := json.Marshal(db)
  if err != nil {
    return err
  }
  err = os.WriteFile(fileName, encodedData, 0666)
  if err != nil {
    return err
  }
  return nil
}

func AuthAdd(conf Config) error {
  Database, err := openDatabase()
  if err != nil {
    return err
  }

  // Checks if account is present
  if _, exists := Database.Accounts[conf.Username]; exists != false {
    return errors.New("Username already present. If you want to update it use auth-update")
  }
  Database.Accounts[conf.Username] = conf

  err = writeDatabase(Database)
  if err != nil {
    return err
  }
  return nil
}

func AuthDelete(username string) error {
  Database, err := openDatabase()
  if err != nil {
    return err
  }

  delete(Database.Accounts, username)

  err = writeDatabase(Database)
  if err != nil {
    return err
  }
  return nil
}
