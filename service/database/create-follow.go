package database

/**
* Adds new following to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateFollow(user uint64, follow uint64) error {
	// check if users exist
	var username string
	row1 := db.c.QueryRow(`SELECT USERNAME FROM USERS WHERE USER-ID=?`, user)
	if row1.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row2 := db.c.QueryRow(`SELECT USERNAME FROM USERS WHERE USER-ID=?`, follow)
	if row2.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user is banned
	row := db.c.QueryRow(`SELECT USER-ID FROM BAN WHERE USER-ID=? AND BAN-ID=?`, follow, user)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see following's profile
	if err == nil {
		return ErrUserBanned
	}

	// insert follow in database
	_, err = db.c.Exec(`INSERT INTO FOLLOWING (USER-ID, FOLLOW-ID) VALUES (?, ?)`, user, follow)
	if err != nil {
		return err
	}

	// even if already followed no error returned
	return nil
}
