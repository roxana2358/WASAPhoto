package database

/**
* Removes follow from database; returns error if the request is unsuccessfull.
*/
func (db *appdbimpl) DeleteFollow(user uint64, unfollow uint64) error {
	// delete follow from database
	res, err := db.c.Exec(`DELETE FROM FOLLOWING (USER-ID, FOLLOW-ID) VALUES (?, ?)`, user, unfollow)
	if err != nil {
		return err
	}
	
	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrFollowNotFound
	}
	return nil
}
