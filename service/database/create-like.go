package database

/**
* Adds new like to database; returns an error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateLike(userID uint64, postID uint64) error {
	// get the id of the owner of the photo
	row := db.c.QueryRow(`SELECT UserId FROM POSTS WHERE PostId=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the id, the post is not in database
		return ErrPostNotFound
	}

	// check if user is banned
	row = db.c.QueryRow(`SELECT BannedId FROM BAN WHERE UserId=? AND BannedId=?`, owner, userID)
	var banned uint64
	if row.Scan(&banned) == nil {
		// if there is a row the user was banned
		return ErrUserBanned
	}

	// like photo
	_, err := db.c.Exec(`INSERT INTO LIKES (PostId, UserId) VALUES (?, ?)`, postID, userID)
	if err != nil {
		return err
	}

	return nil
}
