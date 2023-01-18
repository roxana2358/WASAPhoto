package database

import "errors"

var ErrUserNotFound = errors.New("user not in database")
var ErrUserBanned = errors.New("information cannot be fetched")
var ErrUsernameNotAvailable = errors.New("this username is not available")
var ErrFollowNotFound = errors.New("currently not following this user")
var ErrBanNotFound = errors.New("user currently not banned")
var ErrPostNotFound = errors.New("post not in database")
var ErrFileNotFound = errors.New("file not found in database")
var ErrUnauthorized = errors.New("the user is not authorized for this action")
var ErrCommentNotFound = errors.New("the comment is not in database")
var ErrCouldNotFollow = errors.New("you should unban the user before following")

type Userprofile struct {
	Id             uint64
	Username       string
	Photos         []uint64
	NumberOfPhotos int
	Followers      []string
	Following      []string
	Banned         []string
}

type Userpost struct {
	UserID   uint64
	Username string
	PostID   uint64
	Date     string
	Time     string
	Likes    int
	Comments int
}

type postList []Userpost

func (p postList) Len() int {
	return len(p)
}

func (p postList) Less(i, j int) bool {
	return p[i].Date > p[j].Date || p[i].Time > p[j].Time
}

func (p postList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
