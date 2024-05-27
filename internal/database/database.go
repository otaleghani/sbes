package database

import (
	"encoding/json"
	"os"
)

type Account struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	SmtpHost     string `json:"smtpHost"`
	SmtpPort     int    `json:"smtpPort"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type OAuthClient struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Message struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	MsgType string `json:"msg_type"` // either plain or html
}

type MailingList struct {
	List []string `json:"list"`
}

type Database struct {
	Accounts     map[string]Account     `json:"accounts"`
	OAuthClients map[string]OAuthClient `json:"oauths_clients"`
	MailingLists map[string]MailingList `json:"mailing_lists"`
	Messages     map[string]Message     `json:"messages"`
}

func databasePath() (string, string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	path := homePath + "/.cache/sbes/"
	fileName := path + "db.json"

	return path, fileName, nil
}

func openDatabase() (Database, error) {
	path, fileName, err := databasePath()
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
	Db := Database{
		Accounts:     make(map[string]Account),
		OAuthClients: make(map[string]OAuthClient),
		MailingLists: make(map[string]MailingList),
		Messages:     make(map[string]Message),
	}
	if err = json.Unmarshal(data, &Db); err != nil {
		return Database{}, err
	}
	return Db, nil
}

func writeDatabase(db Database) error {
	_, fileName, err := databasePath()
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
