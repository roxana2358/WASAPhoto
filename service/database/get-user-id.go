package database

/**
* Cheks if username in database: if present returns the id and no error, otherwise returns (nil, error).
 */
func (db *appdbimpl) GetUserId(username string) (uint64, error) {
	// QUERY: find the id of requested username
	row := db.c.QueryRow(`SELECT ID FROM USERS WHERE USERNAME=?`, username)

	var id uint64
	// if there is no row with the username, the user is not in database
	if row.Scan(&id) != nil {
		return 0, ErrUserNotFound
	}
	return id, nil
}