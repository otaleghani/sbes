// Helper functions to work with the database

package database

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
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
  Name string `json:"name"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	MsgType string `json:"msg_type"` // either plain or html
}

type MailingList struct {
  Name string `json:"name"`
	List []string `json:"list"`
}

type TrackedOpen struct {
	Recipient string    `json:"recipient"`
	Data      time.Time `json:"data"`
}

type TrackedClick struct {
	Recipient string    `json:"recipient"`
	Link      string    `json:"link"`
	Data      time.Time `json:"data"`
}

type Campaign struct {
	Name          string                  `json:"name"`
	From          Account                 `json:"from"`
	Msg           Message                 `json:"msg"`
	To            MailingList             `json:"to"`
	TrackedOpens   map[string][]TrackedOpen  `json:"tracked_opens"`
	TrackedClicks map[string][]TrackedClick `json:"tracked_clicks"`
}

type Database struct {
  Domain string `json:"domain"`
	Accounts     map[string]Account     `json:"accounts"`
	OAuthClients map[string]OAuthClient `json:"oauths_clients"`
	MailingLists map[string]MailingList `json:"mailing_lists"`
	Messages     map[string]Message     `json:"messages"`
	Campaigns    map[string]Campaign    `json:"campaigns"`
}

func databasePath() (string, string, error) {
	// Finds home directory
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}

	// Creates $HOME/.cache.sbes/ and $HOME/.cache.sbes/db.json
	path := homePath + "/.cache/sbes/"
	fileName := path + "db.json"

	// Returns path and filename
	return path, fileName, nil
}

func openDatabase() (Database, error) {
	// Takes path and file name
	path, fileName, err := databasePath()
	if err != nil {
		return Database{}, err
	}

	// Clean the input path
	cleanPath := filepath.Clean(fileName)

	baseDir := "/your/base/directory" // Change to your base directory
	if !filepath.IsAbs(cleanPath) {
		cleanPath = filepath.Join(baseDir, cleanPath)
	}

	// Tests if the file is present
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		// Creates all the paths
		err = os.MkdirAll(path, 0750)
		if err != nil {
			return Database{}, err
		}
		// Touches file
		err = os.WriteFile(cleanPath, []byte("{}"), 0600)
		if err != nil {
			return Database{}, err
		}
	}

	// Parses data from file
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return Database{}, err
	}

	// Marshals the json
	Db := Database{
		Accounts:     make(map[string]Account),
		OAuthClients: make(map[string]OAuthClient),
		MailingLists: make(map[string]MailingList),
		Messages:     make(map[string]Message),
    Campaigns:    make(map[string]Campaign),
	}
	if err = json.Unmarshal(data, &Db); err != nil {
		return Database{}, err
	}

	// Returns the database struct
	return Db, nil
}

func writeDatabase(db Database) error {
	// Takes path and file name
	_, fileName, err := databasePath()
	if err != nil {
		return err
	}

	// Encodes json
	encodedData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	// Writes data in file
	err = os.WriteFile(fileName, encodedData, 0600)
	if err != nil {
		return err
	}
	return nil
}
