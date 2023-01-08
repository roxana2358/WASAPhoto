package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"wasa-photo/service/database"
)

var ErrUnauthorized = errors.New("user not in database")

type Username struct {
	Username string //`json:"username"`
}

type ID struct {
	Id uint64 //`json:"ID"`
}

type Comment struct {
	Comment string //`json:"comment"`
}

type Userprofile struct {
	Username       string   //`json:"username"`
	Photos         []uint64 //`json:"photos"`
	NumberOfPhotos int      //`json:"numberOfPhotos"`
	Followers      []string //`json:"followers"`
	Following      []string //`json:"following"`
}

type Userpost struct {
	Username string //`json:"username"`
	PostID   uint64 //`json:"photo"`
	Date     string //`json:"date"`
	Time     string //`json:"time"`
	Likes    int    //`json:"likes"`
	Comments int    //`json:"comments"`
}

/**
* Gets token from header.
 */
func getHeaderToken(r *http.Request) (uint64, error) {
	auth := strings.Split(r.Header.Get("Authorization"), " ")[0]
	// token absent
	if strings.Compare(auth, "") == 0 {
		return 0, ErrUnauthorized
	}
	token, err := strconv.ParseUint(auth, 10, 64)
	// token extraction failed
	if err != nil {
		return 0, err
	}
	return token, nil
}

/**
* Checks if the user that requested the action is the same as the one that is supposed to do it.
 */
func checkAuth(srcUser uint64, dstUser uint64) bool {
	return srcUser == dstUser
}

/**
* UserProfileFromDatabase populates the struct with data from the database, overwriting all values.
 */
func (u *Userprofile) UserProfileFromDatabase(userprofile database.Userprofile) {
	u.Username = userprofile.Username
	u.Photos = userprofile.Photos
	u.NumberOfPhotos = userprofile.NumberOfPhotos
	u.Followers = userprofile.Followers
	u.Following = userprofile.Following
}

/**
* UserPostFromDatabase populates the struct with data from the database, overwriting all values.
 */
func (u *Userpost) UserPostFromDatabase(userpost database.Userpost) {
	u.Username = userpost.Username
	u.PostID = userpost.PostID
	u.Date = userpost.Date
	u.Time = userpost.Time
	u.Likes = userpost.Likes
	u.Comments = userpost.Comments
}
