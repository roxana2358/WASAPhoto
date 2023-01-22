package database

/**
* Adds new photo to database; returns id or error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreatePhoto(userID uint64, postId uint64, time string, date string, fileName string) (uint64, error) {
	// INSERT photo in database
	_, err := db.c.Exec(`INSERT INTO Posts (PostId, UserId, Filename, Time, Date) VALUES (?, ?, ?, ?, ?)`, postId, userID, fileName, time, date)
	if err != nil {
		return 0, err
	}

	// return id and no error
	return postId, nil
}
