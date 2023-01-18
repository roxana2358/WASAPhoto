package database

/**
* Removes follow from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeleteFollow(userID uint64, unfollowID uint64) error {
	// delete follow from database
	res, err := db.c.Exec(`DELETE FROM Following WHERE UserId=? AND FollowingId=?`, userID, unfollowID)
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
