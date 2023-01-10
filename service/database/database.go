/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"image"
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
	CreatePhoto(user uint64, time string, date string, image image.Image) (uint64, error)
	GetImage(imageId uint64) (image.Image, error)
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

	_, err := db.Exec("PRAGMA foreign_keys = ON; ")
	if err != nil {
		return nil, err
	}

	// Check if tables exist. If not, the database is empty, and we need to create the structure
	var tableName string

	// USERS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='USERS';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE USERS 
						(Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
						 Username TEXT PRIMARY KEY
				 		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// FOLLOWING (UserId FOLLOWES FollowingID)
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='FOLLOWING';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE FOLLOWING 
						(UserId INTEGER NOT NULL, 
						 FollowingID INTEGER NOT NULL,
						 FOREIGN KEY(UserId) REFERENCES USERS(Id),
						 FOREIGN KEY(FollowingID) REFERENCES USERS(Id)
				 		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// BAN (UserId BANNED BannedId)
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='BAN';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE BAN 
						(UserId INTEGER NOT NULL, 
						 BannedId INTEGER NOT NULL,
						 FOREIGN KEY(UserId) REFERENCES USERS(Id),
						 FOREIGN KEY(BannedId) REFERENCES USERS(Id)
				 		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// POSTS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='POSTS';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE POSTS 
						(PostId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
						 UserId INTEGER,
						 Filename TEXT,
						 Time TEXT,
						 Date TEXT,
						 FOREIGN KEY(UserId) REFERENCES USERS(Id)
				 		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// LIKES
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='LIKES';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE LIKES 
						(PostId INTEGER NOT NULL,
						 UserId INTEGER NOT NULL,
						 FOREIGN KEY(PostId) REFERENCES POSTS(PostId),
						 FOREIGN KEY(UserId) REFERENCES USERS(Id)
				 		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// COMMENTS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='COMMENTS';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE COMMENTS 
						(CommentId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						 PostId INTEGER NOT NULL,
						 UserId INTEGER NOT NULL, 
						 Text TEXT,
						 FOREIGN KEY(PostId) REFERENCES POSTS(PostId),
						 FOREIGN KEY(UserId) REFERENCES USERS(Id)
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
