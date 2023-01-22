package database

/**
* Gets the filename of the image with imageId.
 */
func (db *appdbimpl) GetNextPostId() uint64 {
	//get next id
	row := db.c.QueryRow(`SELECT PostId FROM Posts WHERE PostId=(SELECT max(PostId) FROM Posts)`)
	var postId uint64
	if row.Scan(&postId) != nil {
		postId = 0
	}
	postId++
	return postId
}
