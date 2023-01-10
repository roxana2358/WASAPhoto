package database

/**
* Adds new ban to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateBan(userID uint64, banID uint64) error {
	// check if users exist
	var username string
	row1 := db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, userID)
	if row1.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row2 := db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, banID)
	if row2.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user was banned
	row := db.c.QueryRow(`SELECT UserId FROM Ban WHERE UserId=? AND BannedId=?`, banID, userID)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see ban's profile
	if err == nil {
		return ErrUserBanned
	}

	// insert ban in database
	_, err = db.c.Exec(`INSERT INTO Ban (UserId, BannedId) VALUES (?, ?)`, userID, banID)
	if err != nil {
		return err
	}

	return nil
}
