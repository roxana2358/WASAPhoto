package database

import (
	"database/sql"
	"errors"
)

/**
* Removes image from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeletePhoto(userID uint64, postID uint64) (string, error) {
	// see if user is the owner of photo
	row := db.c.QueryRow(`SELECT UserId FROM Posts WHERE PostId=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the postID, the photo is not in database
		return "", ErrPostNotFound
	}
	if owner != userID {
		// the user that requested the action is not the owner of the photo
		return "", ErrUnauthorized
	}

	// get filename
	row = db.c.QueryRow(`SELECT Filename FROM Posts WHERE PostId=?`, postID)
	var filename string
	if row.Scan(&filename) != nil {
		// if there is no row with the postID, the photo is not in database
		return "", ErrPostNotFound
	}

	// delete all likes on post
	_, err := db.c.Exec(`DELETE FROM Likes WHERE PostId=?`, postID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	// delete all comments on post
	_, err = db.c.Exec(`DELETE FROM Comments WHERE PostId=?`, postID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	// delete photo from database
	res, err := db.c.Exec(`DELETE FROM Posts WHERE PostId=?`, postID)
	if err != nil {
		return "", err
	}

	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if affected == 0 {
		return "", ErrPostNotFound
	}

	return filename, nil
}
