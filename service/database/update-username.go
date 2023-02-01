package database

/**
* Updates username of user with given id; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) UpdateUsername(id uint64, newUsername string) error {
	// check if username already taken
	row := db.c.QueryRow(`SELECT Id FROM Users WHERE Username=?`, newUsername)
	var otherId uint64
	if row.Scan(&otherId) == nil && id != otherId {
		// username not available
		return ErrUsernameNotAvailable
	}

	// if not, try updating username
	res, err := db.c.Exec(`UPDATE Users SET Username=? WHERE Id=?`, newUsername, id)
	if err != nil {
		return err
	}

	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// if it didn't update any username, then the user didn't exist
		return ErrUserNotFound
	}
	return nil
}
