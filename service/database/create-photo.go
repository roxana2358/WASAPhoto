package database

import (
	"image"
	"image/png"
	"os"
	"strconv"
)

/**
* Adds new photo to database; returns id or error if the request is unsuccessfull.
 */
func (db *appdbimpl) CreatePhoto(user uint64, time string, date string, image image.Image) (uint64, error) {
	// INSERT photo in database
	res, err := db.c.Exec(`INSERT INTO POSTS (USER-ID, TIME, DATE) VALUES (?, ?, ?)`, user, time, date)
	if err != nil {
		return 0, err
	}

	// check if it affected the DB
	_, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	// save photo locally and add info to db
	fileName := "./" + strconv.FormatUint(user, 64) + "-" + time + "-" + date
	imgFile, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer imgFile.Close()
	png.Encode(imgFile, image)
	res, err = db.c.Exec(`UPDATE POSTS SET FILENAME=? WHERE USER-ID=?`, fileName, user)
	if err != nil {
		return 0, err
	}

	// get its id
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// return id and no error
	return uint64(lastInsertID), nil
}
