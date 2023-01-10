package database

/**
* Adds new user to database and returns (id, nil) if no errors occured.
 */
func (db *appdbimpl) CreateUser(username string) (uint64, error) {
	// INSERT user in database
	res, err := db.c.Exec(`INSERT INTO USERS (USERNAME) VALUES (?)`, username)
	if err != nil {
		return 0, err
	}

	// get its id
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
