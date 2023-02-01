package database

/**
* Removes comment from database; returns error if the request is unsuccessfull.
 */
func (db *appdbimpl) DeleteComment(userID uint64, postID uint64, commentID uint64) error {
	// see if user is the owner of comment
	row := db.c.QueryRow(`SELECT UserId FROM Comments WHERE PostId=? AND CommentId=?`, postID, commentID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the commentID, the comment is not in database
		return ErrCommentNotFound
	}
	if owner != userID {
		// the user that requested the action is not the owner of the comment
		return ErrUnauthorized
	}

	// delete comment from database
	res, err := db.c.Exec(`DELETE FROM Comments WHERE CommentId=?`, commentID)
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
