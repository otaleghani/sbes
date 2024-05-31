package database

import (
	"errors"
	"fmt"
	"time"
)

func CampaignAdd(name, account, msg, to string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Checks if user already exist
	if _, exists := Db.Campaigns[name]; exists {
		return errors.New("campaign name already present")
	}

	accountObj, err := AccountGetObject(account)
	if err != nil {
		return err
	}

	msgObj, err := MessageGetObject(msg)
	if err != nil {
		return err
	}

	mlistObj, err := MailingListGetObject(to)
	if err != nil {
		return err
	}

	Db.Campaigns[name] = Campaign{
		Name:          name,
		From:          accountObj,
		Msg:           msgObj,
		To:            mlistObj,
		TrackedOpens:  make(map[string][]TrackedOpen),
		TrackedClicks: make(map[string][]TrackedClick),
	}

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func TrackedOpenAdd(campaignName, email string, time time.Time) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	campaign, exists := Db.Campaigns[campaignName]
	if !exists {
		return errors.New("campaign does not exist")
	}

	_, exists = campaign.TrackedOpens[email]
	if !exists {
		Db.Campaigns[campaignName].TrackedOpens[email] = []TrackedOpen{}
	}
	tOpens := campaign.TrackedOpens[email]
	tOpens = append(tOpens, TrackedOpen{Recipient: email, Data: time})
	Db.Campaigns[campaignName].TrackedOpens[email] = tOpens

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func TrackedClickAdd(campaignName, email, link string, time time.Time) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	campaign, exists := Db.Campaigns[campaignName]
	if !exists {
		return errors.New("campaign does not exist")
	}

	_, exists = campaign.TrackedClicks[email]
	if !exists {
		Db.Campaigns[campaignName].TrackedClicks[email] = []TrackedClick{}
	}
	tClicks := campaign.TrackedClicks[email]
	tClicks = append(tClicks, TrackedClick{
		Recipient: email,
		Link:      link,
		Data:      time,
	})
	Db.Campaigns[campaignName].TrackedClicks[email] = tClicks

	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func CampaignDelete(campaignName string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	// Deletes account
	delete(Db.Campaigns, campaignName)

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func CampaignsList() error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	fmt.Printf(
		"%-20v %-20v %16v %16v\n",
		"NAME", "ACCOUNT", "MESSAGE", "LIST")
	for name, val := range Db.Campaigns {
		truncName := name
		if len(truncName) > 19 {
			truncName = truncName[:19]
		}
		truncAccount := val.From.Username
		if len(truncAccount) > 19 {
			truncAccount = truncAccount[:19]
		}
		truncMessage := val.Msg.Name
		if len(truncMessage) > 19 {
			truncMessage = truncMessage[:19]
		}
		truncMailingList := val.To.Name
		if len(truncMailingList) > 19 {
			truncMailingList = truncMailingList[:19]
		}
		fmt.Printf(
			"%-20v %-20v %16v %16v\n",
			truncName, truncAccount, truncMessage, truncMailingList)
	}
	return nil
}


func CampaignGet(name string) (
	string, Account, Message, MailingList, map[string][]TrackedOpen,
	map[string][]TrackedClick, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", Account{}, Message{}, MailingList{}, nil, nil, err
	}

	// Checks if record exists
	val, exists := Db.Campaigns[name]
	if !exists {
		return "", Account{}, Message{}, MailingList{}, nil, nil, errors.New("record does not exist")
	}

	return val.Name, val.From, val.Msg, val.To, val.TrackedOpens, val.TrackedClicks, nil
}
