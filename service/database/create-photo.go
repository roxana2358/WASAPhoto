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
func (db *appdbimpl) CreatePhoto(userID uint64, time string, date string, image image.Image) (uint64, error) {
	// INSERT photo in database
	res, err := db.c.Exec(`INSERT INTO POSTS (UserId, Time, Date) VALUES (?, ?, ?)`, userID, time, date)
	if err != nil {
		return 0, err
	}

	// get its id
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	// save photo locally and add info to db
	fileName := "./photos/" + strconv.FormatInt(lastInsertID, 10) + ".png"
	imgFile, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer imgFile.Close()
	err = png.Encode(imgFile, image)
	if err != nil {
		return 0, err
	}
	_, err = db.c.Exec(`UPDATE POSTS SET Filename=? WHERE PostId=?`, fileName, lastInsertID)
	if err != nil {
		return 0, err
	}

	// return id and no error
	return uint64(lastInsertID), nil
}
