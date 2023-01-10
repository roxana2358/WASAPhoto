package database

/**
* Removes comment from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeleteComment(userID uint64, postID uint64, commentID uint64) error {
	// delete comment from database
	res, err := db.c.Exec(`DELETE FROM COMMENTS WHERE PostId=? AND CommentId=?`, postID, commentID)
	if err != nil {
		return err
	}

	// check if it affected the DB
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrCommentNotFound
	}
	return nil
}
