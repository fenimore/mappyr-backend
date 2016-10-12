package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
)

const (
	DB_USER     = "mappyr"
	DB_PASSWORD = "mappass"
	DB_NAME     = "mappyr"
)

/* Database Helpers */
// InitDB Opens a new sqlite3 db in path
func InitDB() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}

// CreateTable creates the database tables
func CreateTable(db *sql.DB) error {
	// TODO add binary for photographs
	// TODO: add place name?
	// user field is for related user id
	comment_schema := `
CREATE TABLE IF NOT EXISTS comments(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    upvotes INTEGER DEFAULT 0,
    downvotes INTEGER DEFAULT 0,
    date DATETIME,
    user INTEGER DEFAULT 0
);
`
	// DON't forget to hash password
	user_schema := `
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    password TEXT,
    date DATETIME
);
`
	_, err := db.Exec(user_schema)
	if err != nil {
		return err
	}
	_, err = db.Exec(comment_schema)
	if err != nil {
		return err
	}
	return nil
}

// MockComment inserts a fake comment for testing
// Only Works for SQLITE
// DEPRECATED
func MockComment(db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO comments(title, " +
		"description, lat, lon," +
		"date, user)VALUES($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec("Great food!", "Although crap service",
		"41.353", "-71.113", time.Now(), 0)
	if err != nil {
		return -1, err
	}
	//id, err := res.LastInsertId()
	//if err != nil {
	//	return -1, err
	//}
	return id, nil
}

/* DB read */

// ReadComment reads a comment from the datase with an id.
func ReadComment(db *sql.DB, id int) (Comment, error) {
	rows, err := db.Query("select * from comments where id = $1", id)
	c := Comment{}
	if err != nil {
		return c, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&c.Id, &c.Title,
			&c.Description,
			&c.Lat, &c.Lon,
			&c.Upvotes, &c.Downvotes,
			&c.Date, &c.UserId)
	}
	rows.Close()
	if c.Id == 0 {
		return c, errors.New("Id does not exist")
	}
	return c, nil
}

// ReadComments returns all comments
func ReadComments(db *sql.DB) ([]Comment, error) {
	comments := make([]Comment, 0)

	stmt := "SELECT * FROM comments"
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := Comment{}
		err = rows.Scan(&c.Id, &c.Title,
			&c.Description,
			&c.Lat, &c.Lon,
			&c.Upvotes, &c.Downvotes,
			&c.Date, &c.UserId)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	rows.Close() // Redundant is good
	return comments, nil
}

/* DB Write */
// WriteComment
func WriteComment(db *sql.DB, c Comment) (int64, error) {
	var lastInsertId int
	err = db.QueryRow("INSERT INTO comments(title,description,lat, lon,date,user) VALUES($1,$2,$3,$4,$5) returning id;", c.Title, c.Description, c.Lat, c.Lon, time.Now(), c.UserId).Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}
	return lastInsertId, nil
}

/* Update DB */
// UpVoteComment
func UpVoteComment(db *sql.DB, id int) error {
	stmt, err := db.Prepare("UPDATE comments SET upvotes = upvotes + 1 where id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// DownVoteComment downvotes a row
func DownVoteComment(db *sql.DB, id int) error {
	stmt, err := db.Prepare("UPDATE comments SET downvotes = downvotes + 1 where id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

/* Delete */
func Delete(db *sql.DB, id int) error {
	stmt, err := db.Prepare("delete FROM comments WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
