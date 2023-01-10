package database

/**
* Gets the stream with 30 post from following in reverse chronological order.
 */
func (db *appdbimpl) GetUserStream(userID uint64) ([]Userpost, error) {
	var userStream []Userpost
	var userPost Userpost

	// request 30 posts
	rows, err := db.c.Query(`SELECT USERS.Username, POSTS.PostId, POSTS.Date, POSTS.Time
							FROM FOLLOWING 
							INNER JOIN POSTS ON FOLLOWING.FollowingID=POSTS.UserId
							INNER JOIN USERS ON FOLLOWING.FollowingID=USERS.Id
							WHERE FOLLOWING.UserId=?
							LIMIT 30
							ORDERBY Date DESC,Time DESC`, userID)
	if err != nil {
		return userStream, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be put in userStream
	for rows.Next() {
		err = rows.Scan(&userPost.Username, &userPost.PostID, &userPost.Date, &userPost.Time)
		if err != nil {
			return userStream, err
		}
		// get number of likes
		row := db.c.QueryRow(`SELECT count(*) 
							FROM POSTS 
							INNER JOIN LIKES ON POSTS.PostId=LIKES.PostId 
							WHERE PostId=?`, userPost.PostID)
		var likes int
		err = row.Scan(&likes)
		if err != nil {
			return userStream, err
		}
		userPost.Likes = likes
		// get number of comments
		row = db.c.QueryRow(`SELECT count(*) 
							FROM POSTS 
							INNER JOIN COMMENTS ON POSTS.PostId=COMMENTS.PostId 
							WHERE PostId=?`, userPost.PostID)
		var comments int
		err = row.Scan(&comments)
		if err != nil {
			return userStream, err
		}
		userPost.Comments = comments
		// add userPost to userStream
		userStream = append(userStream, userPost)
	}
	if err = rows.Err(); err != nil {
		return userStream, err
	}
	return userStream, nil
}
