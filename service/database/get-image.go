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
	row := db.c.QueryRow(`SELECT FILENAME FROM POSTS WHERE POST-ID=?`, imageId)
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
	if err != nil {
		return nil, ErrFileNotFound
	}
	// decode
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, ErrDecode
	}

	// return image
	return img, nil
}