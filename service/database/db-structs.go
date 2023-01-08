package database

import "errors"

var ErrUserNotFound = errors.New("user not in database")
var ErrUserBanned = errors.New("information cannot be fetched")
var ErrUsernameNotAvailable = errors.New("this username is not available")
var ErrFollowNotFound = errors.New("currently not following this user")
var ErrBanNotFound = errors.New("user currently not banned")
var ErrPostNotFound = errors.New("post not in database")
var ErrFileNotFound = errors.New("file not found in database")
var ErrDecode = errors.New("decoding error")
var ErrUnauthorized = errors.New("the user is not authorized for this action")
var ErrCommentNotFound = errors.New("the comment is not in database")

type Userprofile struct {
	Username       string
	Photos         []uint64
	NumberOfPhotos int
	Followers      []string
	Following      []string
}

type Userpost struct {
	Username string
	PostID   uint64
	Date     string
	Time     string
	Likes    int
	Comments int
}
