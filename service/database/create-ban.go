package database

/**
* Adds new ban to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateBan(userID uint64, banID uint64) error {
	// check if users exist
	var username string
	row1 := db.c.QueryRow(`SELECT USERNAME FROM USERS WHERE ID=?`, userID)
	if row1.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row2 := db.c.QueryRow(`SELECT USERNAME FROM USERS WHERE ID=?`, banID)
	if row2.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user was banned
	row := db.c.QueryRow(`SELECT USER-ID FROM BAN WHERE USER-ID=? AND BANNED-ID=?`, banID, userID)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see ban's profile
	if err == nil {
		return ErrUserBanned
	}

	// insert ban in database
	_, err = db.c.Exec(`INSERT INTO BAN (USER-ID, BANNED-ID) VALUES (?, ?)`, userID, banID)
	if err != nil {
		return err
	}

	return nil
}
