package database

func UpdateDomainTracker(domain string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	Db.DomainTracker = domain

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func DomainTrackerGet() (string, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", err
	}

	return Db.DomainTracker, nil
}

func UpdateDomainOAuth(domain string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	Db.DomainOAuth = domain

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func DomainOAuthGet() (string, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", err
	}

	return Db.DomainOAuth, nil
}
