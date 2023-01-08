package database

/**
* Adds new ban to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateBan(user uint64, ban uint64) error {
	// check if users exist
	var username string
	row1 := db.c.QueryRow(`SELECT USERNAME FROM USERS WHERE USER-ID=?`, user)
	if row1.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row2 := db.c.QueryRow(`SELECT USERNAME FROM USERS WHERE USER-ID=?`, ban)
	if row2.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user was banned
	row := db.c.QueryRow(`SELECT USER-ID FROM BAN WHERE USER-ID=? AND BAN-ID=?`, ban, user)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see following's profile
	if err == nil {
		return ErrUserBanned
	}

	// insert ban in database
	_, err = db.c.Exec(`INSERT INTO BAN (USER-ID, FOLLOW-ID) VALUES (?, ?)`, user, ban)
	if err != nil {
		return err
	}

	// even if already banned no error returned
	return nil
}
