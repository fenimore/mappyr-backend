package database

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/* Database Helpers */
// InitDB Opens a new sqlite3 db in path
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./mappyr.db")
	if err != nil {
		return nil, err
	}
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
    text TEXT NOT NULL,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
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
func MockComment(db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO comments(text, lat, lon," +
		"date, user)values(?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec("Great food!", "41.333", "-71.333",
		time.Now(), 0)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/* DB read */

// ReadComment reads a comment from the datase with an id.
func ReadComment(db *sql.DB, id int) (Comment, error) {
	rows, err := db.Query("select * from comments where id = ?", id)
	c := Comment{}
	if err != nil {
		return c, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&c.Id, &c.Text,
			&c.Lat, &c.Lon,
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
		err = rows.Scan(&c.Id, &c.Text,
			&c.Lat, &c.Lon,
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
	stmt, err := db.Prepare("INSERT INTO comments(text, lat, lon," +
		"date, user)values(?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(c.Text, c.Lat, c.Lon,
		time.Now(), c.UserId)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Put/ delete/ Create
