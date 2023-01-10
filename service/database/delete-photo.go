package database

/**
* Removes image from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeletePhoto(userID uint64, postID uint64) error {
	// see if user is the owner of photo
	row := db.c.QueryRow(`SELECT USER-ID FROM POSTS WHERE POST-ID=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the postID, the photo is not in database
		return ErrPostNotFound
	}
	if owner != userID {
		// the user that requested the action is not the owner of the photo
		return ErrUnauthorized
	}

	// delete photo from database
	res, err := db.c.Exec(`DELETE FROM POSTS WHERE POST-ID=?`, postID)
	if err != nil {
		return err
	}

	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrPostNotFound
	}

	// delete all likes on post
	_, err = db.c.Exec(`DELETE FROM LIKES WHERE POST-ID=?`, postID)
	if err != nil {
		return err
	}

	// delete all comments on post
	_, err = db.c.Exec(`DELETE FROM COMMENTS WHERE POST-ID=?`, postID)
	if err != nil {
		return err
	}

	return nil
}
