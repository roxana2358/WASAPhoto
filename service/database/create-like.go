package database

/**
* Adds new like to database; returns an error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateLike(userID uint64, postID uint64) error {
	// get the id of the owner of the photo
	row := db.c.QueryRow(`SELECT UserId FROM Posts WHERE PostId=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the id, the post is not in database
		return ErrPostNotFound
	}

	// check if user exists
	var username string
	row = db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, userID)
	if row.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user is banned
	row = db.c.QueryRow(`SELECT BannedId FROM Ban WHERE UserId=? AND BannedId=?`, owner, userID)
	var banned uint64
	if row.Scan(&banned) == nil {
		// if there is a row the user was banned
		return ErrUserBanned
	}

	// check if already in database
	var id uint64
	row = db.c.QueryRow(`SELECT UserId FROM Likes WHERE UserId=? AND PostId=?`, userID, postID)
	if row.Scan(&id) == nil {
		return nil
	}
	// like photo
	_, err := db.c.Exec(`INSERT INTO Likes (PostId, UserId) VALUES (?, ?)`, postID, userID)
	if err != nil {
		return err
	}

	return nil
}
