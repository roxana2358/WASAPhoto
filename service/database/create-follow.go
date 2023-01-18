package database

/**
* Adds new follow to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateFollow(userID uint64, followID uint64) error {
	// check if users exist
	var username string
	row := db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, userID)
	if row.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row = db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, followID)
	if row.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user is banned
	row = db.c.QueryRow(`SELECT UserId FROM Ban WHERE UserId=? AND BannedId=?`, followID, userID)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see follow's profile
	if err == nil {
		return ErrUserBanned
	}

	// if followID was previously banned userID shouldn't be able to follow
	row = db.c.QueryRow(`SELECT UserId FROM Ban WHERE UserId=? AND BannedId=?`, userID, followID)
	err = row.Scan(&id)
	if err == nil {
		return ErrCouldNotFollow
	}

	// check if already in database
	row = db.c.QueryRow(`SELECT UserId FROM Following WHERE UserId=? AND FollowingId=?`, userID, followID)
	if row.Scan(&id) == nil {
		return nil
	}
	// insert follow in database
	_, err = db.c.Exec(`INSERT INTO Following (UserId, FollowingId) VALUES (?, ?)`, userID, followID)
	if err != nil {
		return err
	}

	return nil
}
