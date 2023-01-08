package database

/**
* Removes image from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeletePhoto(user uint64, post uint64) error {
	// see if user is the owner of photo
	row := db.c.QueryRow(`SELECT USER-ID FROM POSTS WHERE POST-ID=?`, post)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the postId, the photo is not in database
		return ErrFileNotFound
	}
	if owner != user {
		// the user that requested the action is not the owner of the photo
		return ErrUnauthorized
	}

	// delete photo from database
	res, err := db.c.Exec(`DELETE FROM POSTS WHERE POST-ID=?`, post)
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

	return nil
}
