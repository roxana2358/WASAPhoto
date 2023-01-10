package database

import (
	"image"
	"os"
)

/**
* Gets the image from local storage.
 */
func (db *appdbimpl) GetImage(imageId uint64) (image.Image, error) {
	// QUERY: find name of file with imageId
	row := db.c.QueryRow(`SELECT Filename FROM POSTS WHERE PostId=?`, imageId)
	var filename string
	if row.Scan(&filename) != nil {
		return nil, ErrFileNotFound
	}

	// retreive file
	imgFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	// decode
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return img, nil
}
