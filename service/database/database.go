/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserId(username string) (uint64, error)
	CreateUser(username string) (uint64, error)
	GetUserProfile(token uint64, userID uint64) (Userprofile, error)
	UpdateUsername(id uint64, newUsername string) error
	GetUserStream(userID uint64) ([]Userpost, error)
	CreateFollow(user uint64, follow uint64) error
	DeleteFollow(user uint64, unfollow uint64) error
	CreateBan(user uint64, ban uint64) error
	DeleteBan(user uint64, unban uint64) error
	GetNextPostId() (uint64, error)
	CreatePhoto(userID uint64, postId uint64, time string, date string, fileName string) (uint64, error)
	GetImage(imageId uint64) (string, error)
	DeletePhoto(user uint64, post uint64) error
	CreateComment(user uint64, post uint64, comment string) (uint64, error)
	DeleteComment(userID uint64, postID uint64, commentID uint64) error
	CreateLike(userID uint64, postID uint64) error
	DeleteLike(userID uint64, postID uint64) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	// Check if tables exist. If not, the database is empty, and we need to create the structure
	var tableName string

	// Users
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Users (
			Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
			Username TEXT UNIQUE
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// Following (UserId FOLLOWES FollowingId)
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Following';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Following (
			UserId INTEGER NOT NULL, 
			FollowingId INTEGER NOT NULL,
			FOREIGN KEY(UserId) REFERENCES Users(Id),
			FOREIGN KEY(FollowingId) REFERENCES Users(Id),
			PRIMARY KEY(UserId, FollowingId)
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// Ban (UserId BANNED BannedId)
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Ban';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Ban (
			UserId INTEGER NOT NULL, 
			BannedId INTEGER NOT NULL,
			FOREIGN KEY(UserId) REFERENCES Users(Id),
			FOREIGN KEY(BannedId) REFERENCES Users(Id),
			PRIMARY KEY(UserId, BannedId)
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// Posts
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Posts';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Posts (
			PostId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
			UserId INTEGER,
			Filename TEXT,
			Time TEXT,
			Date TEXT,
			FOREIGN KEY(UserId) REFERENCES Users(Id)
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// Likes
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Likes (
			PostId INTEGER NOT NULL,
			UserId INTEGER NOT NULL,
			FOREIGN KEY(PostId) REFERENCES Posts(PostId),
			FOREIGN KEY(UserId) REFERENCES Users(Id),
			PRIMARY KEY(PostId, UserId)
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// Comments
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Comments (
			CommentId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			PostId INTEGER NOT NULL,
			UserId INTEGER NOT NULL, 
			Text TEXT,
			FOREIGN KEY(PostId) REFERENCES Posts(PostId),
			FOREIGN KEY(UserId) REFERENCES Users(Id)
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
