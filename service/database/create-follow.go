package database

/**
* Adds new follow to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateFollow(userID uint64, followID uint64) error {
	// check if users exist
	var username string
	row1 := db.c.QueryRow(`SELECT Username FROM USERS WHERE Id=?`, userID)
	if row1.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row2 := db.c.QueryRow(`SELECT Username FROM USERS WHERE Id=?`, followID)
	if row2.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user is banned
	row := db.c.QueryRow(`SELECT UserId FROM BAN WHERE UserId=? AND BannedId=?`, followID, userID)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see follow's profile
	if err == nil {
		return ErrUserBanned
	}

	// insert follow in database
	_, err = db.c.Exec(`INSERT INTO FOLLOWING (UserId, FollowingID) VALUES (?, ?)`, userID, followID)
	if err != nil {
		return err
	}

	return nil
}
