package database

/**
* Updates username of user with given id; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) UpdateUsername(id uint64, newUsername string) error {
	// check if username already taken (no use of triggers)
	row := db.c.QueryRow(`SELECT ID FROM USERS WHERE USERNAME=?`, newUsername)
	var otherId uint64
	if row.Scan(&otherId) == nil && id != otherId {
		// username not available
		return ErrUsernameNotAvailable
	}

	// if not, try updating usernames
	res, err := db.c.Exec(`UPDATE USERS SET USERNAME=? WHERE ID=?`, newUsername, id)
	if err != nil {
		return err
	}

	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// if we didn't update any username, then the user didn't exist
		return ErrUserNotFound
	}
	return nil
}
