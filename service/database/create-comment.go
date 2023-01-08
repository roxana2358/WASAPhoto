package database

/**
* Adds a new comment in the database and returns (commentID, nil) if no errors occured.
 */
func (db *appdbimpl) CreateComment(userID uint64, postID uint64, comment string) (uint64, error) {
	// get owner id
	row := db.c.QueryRow(`SELECT USER-ID FROM POSTS WHERE POST-ID=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the id, the post is not in database
		return 0, ErrPostNotFound
	}

	// check if user is banned
	row = db.c.QueryRow(`SELECT BANNED-ID FROM BAN WHERE USER-ID=?`, postID)
	var banned uint64
	if row.Scan(&banned) == nil {
		// if there is a row the user was banned
		return 0, ErrUserBanned
	}

	// comment
	res, err := db.c.Exec(`INSERT INTO COMMENTS (POST-ID, USER-ID, TEXT) VALUES (?, ?, ?)`, postID, userID, comment)
	if err != nil {
		return 0, err
	}

	// get comment id
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// return id and no error
	return uint64(lastInsertID), nil
}
