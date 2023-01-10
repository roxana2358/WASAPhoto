package database

/**
* Removes ban from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeleteBan(userID uint64, unbanID uint64) error {
	// delete ban from database
	res, err := db.c.Exec(`DELETE FROM Ban (UserId, BannedId) VALUES (?, ?)`, userID, unbanID)
	if err != nil {
		return err
	}

	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrBanNotFound
	}
	return nil
}
