package database

func UpdateDomain(domain string) error {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return err
	}

	Db.Domain = domain

	// Writes database
	err = writeDatabase(Db)
	if err != nil {
		return err
	}
	return nil
}

func DomainGet() (string, error) {
	// Opens database
	Db, err := openDatabase()
	if err != nil {
		return "", err
	}

	return Db.Domain, nil
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
