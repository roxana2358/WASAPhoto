package database

/**
* Gets the profile of the user with given ID.
 */
func (db *appdbimpl) GetUserProfile(userID uint64, profileID uint64) (Userprofile, error) {
	var userProfile Userprofile
	// check if user is banned
	row := db.c.QueryRow(`SELECT UserId FROM BAN WHERE UserId=? AND BannedId=?`, profileID, userID)
	var id int
	err := row.Scan(&id)
	if err == nil {
		// if row is present token was banned
		return userProfile, ErrUserBanned
	}

	// 1) find username
	row = db.c.QueryRow(`SELECT Username FROM USERS WHERE Id=?`, profileID)
	var username string
	if row.Scan(&username) != nil {
		return userProfile, ErrUserNotFound
	}

	// 2) find photos and their number
	rows, err := db.c.Query(`SELECT PostId FROM POSTS WHERE UserId=?`, profileID)
	if err != nil {
		return userProfile, err
	}
	defer func() { _ = rows.Close() }()
	var photoIds []uint64
	// Here we read the resultset and we build the list to be put in userProfile
	for rows.Next() {
		var photoID uint64
		err = rows.Scan(&photoID)
		if err != nil {
			return userProfile, err
		}
		photoIds = append(photoIds, photoID)
	}
	if err = rows.Err(); err != nil {
		return userProfile, err
	}

	// 3) find following
	rows, err = db.c.Query(`SELECT FollowingID FROM FOLLOWING WHERE UserId=?`, profileID)
	if err != nil {
		return userProfile, err
	}
	defer func() { _ = rows.Close() }()
	var following []string
	var followID uint64
	var follow string
	// Here we read the resultset and we build the list to be put in userProfile
	for rows.Next() {
		err = rows.Scan(&followID)
		if err != nil {
			return userProfile, err
		}

		row = db.c.QueryRow(`SELECT Username FROM USERS WHERE Id=?`, followID)
		if row.Scan(&follow) != nil {
			return userProfile, err
		}

		following = append(following, follow)
	}
	if err = rows.Err(); err != nil {
		return userProfile, err
	}

	// 4) find followers
	rows, err = db.c.Query(`SELECT UserId FROM FOLLOWING WHERE FollowingID=?`, profileID)
	if err != nil {
		return userProfile, err
	}
	defer func() { _ = rows.Close() }()
	var followers []string
	var followerID uint64
	var follower string
	// Here we read the resultset and we build the list to be put in userProfile
	for rows.Next() {
		err = rows.Scan(&followerID)
		if err != nil {
			return userProfile, err
		}

		row = db.c.QueryRow(`SELECT Username FROM USERS WHERE Id=?`, followerID)
		if row.Scan(&follower) != nil {
			return userProfile, err
		}

		followers = append(followers, follower)
	}
	if err = rows.Err(); err != nil {
		return userProfile, err
	}

	// 5) compose userProfile
	userProfile.Username = username
	userProfile.Photos = photoIds
	userProfile.NumberOfPhotos = len(photoIds)
	userProfile.Following = following
	userProfile.Followers = followers

	return userProfile, nil
}
