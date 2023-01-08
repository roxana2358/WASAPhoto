package database

/**
* Removes like from photo and returns an error if any occured.
 */
func (db *appdbimpl) DeleteLike(userID uint64, postID uint64) error {
	// unlike photo
	res, err := db.c.Exec(`DELETE FROM LIKES WHERE POST-ID=? AND USER-ID=?`, postID, userID)
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
