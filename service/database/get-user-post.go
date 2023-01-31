package database

import (
	"database/sql"
	"errors"
)

/**
* Gets the user post.
 */
func (db *appdbimpl) GetUserPost(userID uint64, postID uint64) (Userpost, error) {
	var userPost Userpost
	userPost.PostID = postID

	// get the id of the owner of the post
	row := db.c.QueryRow(`SELECT UserId FROM Posts WHERE PostId=?`, postID)
	var owner uint64
	if row.Scan(&owner) != nil {
		// if there is no row with the id, the post is not in database
		return userPost, ErrPostNotFound
	}

	// check if user is banned
	row = db.c.QueryRow(`SELECT UserId FROM Ban WHERE UserId=? AND BannedId=?`, owner, userID)
	var id int
	err := row.Scan(&id)
	if err == nil {
		// if row is present token was banned
		return userPost, ErrUserBanned
	}

	// request post
	row = db.c.QueryRow(`SELECT Users.Id, Users.Username, Posts.Date, Posts.Time
							FROM Users
							INNER JOIN Posts ON Users.Id=Posts.UserId
							WHERE Posts.PostId=?`, postID)
	err = row.Scan(&userPost.UserID, &userPost.Username, &userPost.Date, &userPost.Time)
	if err != nil {
		return userPost, err
	}

	// get likes
	l, err := db.c.Query(`SELECT Likes.UserId 
							FROM Posts 
							INNER JOIN Likes ON Posts.PostId=Likes.PostId 
							WHERE Posts.PostId=?`, postID)
	if errors.Is(err, sql.ErrNoRows) {
		// no likes
		userPost.Likes = nil
	} else if err == nil {
		// likes
		var likeId uint64
		var likes []uint64
		for l.Next() {
			e := l.Scan(&likeId)
			if e != nil {
				return userPost, err
			}
			likes = append(likes, likeId)
		}
		if err = l.Err(); err != nil {
			return userPost, err
		}
		userPost.Likes = likes
	} else if err != nil {
		// other error
		return userPost, err
	}

	// get comments
	c, err := db.c.Query(`SELECT Users.Username, Comments.UserId, Comments.Text, Comments.CommentId
							FROM Comments
							INNER JOIN Posts ON Posts.PostId=Comments.PostId
							INNER JOIN Users ON Comments.UserId=Users.Id 
							WHERE Comments.PostId=?`, postID)
	if errors.Is(err, sql.ErrNoRows) {
		// no comments
		userPost.Comments = nil
	} else if err == nil {
		// comments
		var comment CommentOBJ
		var comments []CommentOBJ
		for c.Next() {
			err = c.Scan(&comment.Username, &comment.UserID, &comment.Comment, &comment.CommentId)
			if err != nil {
				return userPost, err
			}
			comments = append(comments, comment)
		}
		if err = c.Err(); err != nil {
			return userPost, err
		}
		userPost.Comments = comments
	} else if err != nil {
		// other error
		return userPost, err
	}
	return userPost, nil
}
