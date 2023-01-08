package database

/**
* Adds new like to database; returns id or error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateLike(userID uint64, postID uint64) error {
	// check if user is blocked
	row := db.c.QueryRow(`SELECT BANNED-ID FROM BAN WHERE USER-ID=?`, postID)
	var banned uint64
	if row.Scan(&banned) == nil {
		// if there is a row the user was banned
		return ErrUserBanned
	}

	// like photo
	res, err := db.c.Exec(`INSERT INTO LIKES (POST-ID, USER-ID) VALUES (?, ?)`, postID, userID)
	if err != nil {
		return err
	}

	// check if it affected the database
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrPostNotFound
	}
	return nil
}
