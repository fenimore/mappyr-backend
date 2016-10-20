package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/bmizerany/pq"
	//"github.com/bmizerany/pq" // HEROKU NEEDS FOR PARSING
)

const (
	DB_USER     = "mappyr"  //"kpssujtcjeylzx"
	DB_PASSWORD = "mappass" //"By5bPQQibYr5KDkBu-E9nU5eaO"
	DB_NAME     = "mappyr"  //"dcnd0p0l81l0dr"
	DB_SSL      = "disable" // "require"
)

/* ################################################################################
Database Helpers
   ################################################################################  */

// InitDB Opens a new sqlite3 db in path
func InitDB() (*sql.DB, error) {
	//url := os.Getenv("DATABASE_URL")
	//connection, _ := pq.ParseURL(url) // HEROKU
	// HEROKU: sslmode=require
	// HEROKU: add a space!!
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, DB_NAME, DB_SSL)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	return db, nil
}

// CreateTable creates the database tables
func CreateTable(db *sql.DB) error {
	// TODO add binary for photographs
	// TODO: add place name?
	// user field is for related user id
	comment_schema := `
CREATE TABLE IF NOT EXISTS comments(
    comment_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    upvotes INTEGER DEFAULT 0,
    downvotes INTEGER DEFAULT 0,
    pub_date TIMESTAMP,
    user_id int REFERENCES users (user_id) ON UPDATE CASCADE ON DELETE CASCADE
);`

	// TODO: hash passes and salt dat up
	// TODO: Make username unique
	user_schema := `
CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL,
    password VARCHAR(50),
    create_date TIMESTAMP,
    email VARCHAR(50)
)
`
	vote_schema := `
CREATE TABLE IF NOT EXISTS votes(
    comment_id int REFERENCES comments (comment_id) ON UPDATE CASCADE,
    user_id    int REFERENCES  users    (user_id),
    up         BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT upvote_key PRIMARY KEY (comment_id, user_id)
)
`
	_, err := db.Exec(user_schema)
	if err != nil {
		fmt.Println("user error", err)
		return err
	}

	_, err = db.Exec(comment_schema)
	if err != nil {
		fmt.Println("comments", err)
		return err
	}

	_, err = db.Exec(vote_schema)
	if err != nil {
		fmt.Println("vote", err)
		return err
	}

	return nil
}

/* ################################################################################
Comments
   ################################################################################  */

// ReadComment reads a comment from the datase with an id.
func ReadComment(db *sql.DB, id int) (Comment, error) {
	rows, err := db.Query("select * from comments"+
		" where comment_id = $1", id)
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

// WriteComment
func WriteComment(db *sql.DB, c Comment) (int, error) {
	var lastInsertId int
	err := db.QueryRow("INSERT INTO comments(title,description,lat,lon,pub_date,user_id)"+
		" VALUES($1,$2,$3,$4,$5,$6) returning comment_id;",
		c.Title, c.Description, c.Lat, c.Lon, time.Now(), c.UserId).Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}
	return lastInsertId, nil
}

// UpVoteComment
func VoteComment(db *sql.DB, comment_id, user_id int, up bool) error {
	stmt, err := db.Prepare("INSERT INTO votes(comment_id, user_id, up)VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment_id, user_id, up)
	if err != nil {
		return err
	}

	// UPDATE the comment this effects, so that way we don't have to query
	// on the front end to find how many upvotes downvotes there are
	if up {
		stmt, err := db.Prepare("UPDATE comments SET upvotes = upvotes + 1 " +
			"where comment_id=$1")
		if err != nil {
			return err
		}
		stmt.Exec(comment_id)
	} else {
		stmt, err := db.Prepare("UPDATE comments SET downvotes = downvotes + 1 " +
			"where comment_id=$1")
		if err != nil {
			return err
		}
		stmt.Exec(comment_id)
	}
	return nil
}

// CommentVotes returns a slice of Vote structs according
// to a passed in comment ID
func CommentVotes(comment_id, db *sql.DB) ([]Vote, error) {
	votes := make([]Vote, 0)
	rows, err := db.Query("select * from votes where comment_id = $1", comment_id)
	if err != nil {
		return votes, err
	}
	defer rows.Close()

	for rows.Next() {
		v := Vote{}
		err = rows.Scan(&v.Comment, &v.User, &v.Up)
		if err != nil {
			return votes, err
		}
		votes = append(votes, v)
	}
	rows.Close()
	return votes, nil
}

/* Delete */
func DeleteComment(db *sql.DB, id int) error {
	stmt, err := db.Prepare("delete FROM comments WHERE comment_id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

/* ################################################################################
Users
   ################################################################################  */

func SignUp(db *sql.DB, u User) (int, error) {
	var lastInsertId int
	err := db.QueryRow("INSERT INTO users(user_name, password, create_date, email)"+
		" VALUES($1,$2,$3,$4) returning user_id;",
		u.Name, u.Password, time.Now(),
		u.Email).Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}
	return lastInsertId, nil
}

// ReadUsers returns all comments
func ReadUsers(db *sql.DB) ([]User, error) {
	users := make([]User, 0)

	stmt := "SELECT * FROM users"
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name,
			&u.Password,
			&u.Date, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	rows.Close() // Redundant is good
	return users, nil
}

/* Votes */

func UserVotes(user_id int, db *sql.DB) ([]Vote, error) {
	votes := make([]Vote, 0)
	rows, err := db.Query("select * from votes where user_id = $1", user_id)
	if err != nil {
		return votes, err
	}
	defer rows.Close()

	for rows.Next() {
		v := Vote{}
		err = rows.Scan(&v.Comment, &v.User, &v.Up)
		if err != nil {
			return votes, err
		}
		votes = append(votes, v)
	}
	rows.Close()
	return votes, nil
}

func UserComments(user_id int, db *sql.DB) ([]Comment, error) {
	comments := make([]Comment, 0)
	rows, err := db.Query("select * from comments where user_id = $1", user_id)
	if err != nil {
		return comments, err
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
	rows.Close()
	return comments, nil
}
