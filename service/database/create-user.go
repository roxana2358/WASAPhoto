package database

/**
* Adds new user to database and returns (id, nil) if no errors occured; (nil, error) otherwise.
 */
func (db *appdbimpl) CreateUser(username string) (uint64, error) {
	// INSERT user in database
	res, err := db.c.Exec(`INSERT INTO USERS (USERNAME, POST-NUMBER) VALUES (?, ?)`, username, 0)
	if err != nil {
		return 0, err
	}

	// get its id
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// return id and no error
	return uint64(lastInsertID), nil
}
