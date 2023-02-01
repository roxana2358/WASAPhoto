package database

import (
	"database/sql"
	"errors"
)

/**
* Adds new ban to database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreateBan(userID uint64, banID uint64) error {
	// check if users exist
	var username string
	row := db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, userID)
	if row.Scan(&username) != nil {
		return ErrUserNotFound
	}
	row = db.c.QueryRow(`SELECT Username FROM Users WHERE Id=?`, banID)
	if row.Scan(&username) != nil {
		return ErrUserNotFound
	}

	// check if user was banned
	row = db.c.QueryRow(`SELECT UserId FROM Ban WHERE UserId=? AND BannedId=?`, banID, userID)
	var id int
	err := row.Scan(&id)
	// if row is present user can't see ban's profile
	if err == nil {
		return ErrUserBanned
	}

	// check if already in database
	row = db.c.QueryRow(`SELECT UserId FROM Ban WHERE UserId=? AND BannedId=?`, userID, banID)
	if row.Scan(&id) == nil {
		return nil
	}

	// insert ban in database
	_, err = db.c.Exec(`INSERT INTO Ban (UserId, BannedId) VALUES (?, ?)`, userID, banID)
	if err != nil {
		return err
	}

	// delete follows
	_, err = db.c.Exec(`DELETE FROM Following WHERE UserId=? AND FollowingId=?`, userID, banID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	_, err = db.c.Exec(`DELETE FROM Following WHERE UserId=? AND FollowingId=?`, banID, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return nil
}
