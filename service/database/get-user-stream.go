package database

/**
* Gets the stream with 30 post from following in reverse chronological order.
 */
func (db *appdbimpl) GetUserStream(userID uint64) ([]Userpost, error) {
	var userStream []Userpost
	var userPost Userpost

	// request 30 posts
	rows, err := db.c.Query(`SELECT USERS.USERNAME, POSTS.POST-ID, POSTS.DATE, POSTS.TIME
							FROM FOLLOWING 
							INNER JOIN POSTS ON FOLLOWING.FOLLOWING-ID=POSTS.USER-ID
							INNER JOIN USERS ON FOLLOWING.FOLLOWING-ID=USERS.ID
							WHERE FOLLOWING.USER-ID=?
							LIMIT 30
							ORDERBY DATE DESC,TIME DESC`, userID)
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
							INNER JOIN LIKES ON POSTS.POST-ID=LIKES.POST-ID 
							WHERE POST-ID=?`, userPost.PostID)
		var likes int
		err = row.Scan(&likes)
		if err != nil {
			return userStream, err
		}
		userPost.Likes = likes
		// get number of comments
		row = db.c.QueryRow(`SELECT count(*) 
							FROM POSTS 
							INNER JOIN COMMENTS ON POSTS.POST-ID=COMMENTS.POST-ID 
							WHERE POST-ID=?`, userPost.PostID)
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
