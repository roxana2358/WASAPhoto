package database

/**
* Gets the filename of the image with imageId.
 */
func (db *appdbimpl) GetImage(imageId uint64) (string, error) {
	// QUERY: find name of file with imageId
	row := db.c.QueryRow(`SELECT Filename FROM Posts WHERE PostId=?`, imageId)
	var filename string
	if row.Scan(&filename) != nil {
		return "", ErrPostNotFound
	}

	return filename, nil
}
