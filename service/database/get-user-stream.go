package database

import "sort"

/**
* Gets the stream with 30 post from following in reverse chronological order.
 */
func (db *appdbimpl) GetUserStream(userID uint64) ([]Userpost, error) {
	var userStream []Userpost
	var userPost Userpost

	// request posts
	rows, err := db.c.Query(`SELECT Users.Username, Posts.PostId, Posts.Date, Posts.Time
							FROM Following 
							INNER JOIN Posts ON Following.FollowingID=Posts.UserId
							INNER JOIN Users ON Following.FollowingID=Users.Id
							WHERE Following.UserId=?`, userID)
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
							FROM Posts 
							INNER JOIN Likes ON Posts.PostId=Likes.PostId 
							WHERE PostId=?`, userPost.PostID)
		var likes int
		err = row.Scan(&likes)
		if err != nil {
			return userStream, err
		}
		userPost.Likes = likes
		// get number of comments
		row = db.c.QueryRow(`SELECT count(*) 
							FROM Posts 
							INNER JOIN Comments ON Posts.PostId=Comments.PostId 
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

	// sort and select
	sort.Sort(postList(userStream))
	var finalList []Userpost
	if len(userStream) > 30 {
		finalList = userStream[:30]
	} else {
		finalList = userStream[:]
	}
	return finalList, nil
}
