package database

/**
* Adds a new comment in the database and returns (commentID, nil) if no errors occured.
 */
func (db *appdbimpl) CreateComment(userID uint64, postID uint64, comment string) (uint64, error) {
	// get the id of the owner of the photo
	row := db.c.QueryRow(`SELECT UserId FROM Posts WHERE PostId=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the id, the post is not in database
		return 0, ErrPostNotFound
	}

	// check if user exists
	var username string
	row = db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, userID)
	if row.Scan(&username) != nil {
		return 0, ErrUserNotFound
	}

	// check if user is banned
	row = db.c.QueryRow(`SELECT BannedId FROM Ban WHERE UserId=? AND BannedId=?`, owner, userID)
	var banned uint64
	if row.Scan(&banned) == nil {
		// if there is a row the user was banned
		return 0, ErrUserBanned
	}

	// comment
	res, err := db.c.Exec(`INSERT INTO Comments (PostId, UserId, Text) VALUES (?, ?, ?)`, postID, userID, comment)
	if err != nil {
		return 0, err
	}

	// get comment id
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
