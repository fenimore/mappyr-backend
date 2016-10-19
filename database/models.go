package database

import (
	"fmt"
	"time"

	"database/sql"
	_ "github.com/bmizerany/pq"
)

// TODO: Add a place model?
// TODO:
// Description
// UPvote and Downvote
type Comment struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Lat         float64   `json:"latitude"`
	Lon         float64   `json:"longitude"`
	Upvotes     int       `json:"upvotes"`
	Downvotes   int       `json:"downvotes"`
	Date        time.Time `json:"date"`
	UserId      int       `json:"user-id"` // 0 for anon
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment %i: %s, %d\nAt %f %f",
		c.Id, c.Title, c.Upvotes, c.Lat, c.Lon)
}

type Vote struct {
	Comment int  `json:"comment_id"`
	User    int  `json:"user_id"`
	Up      bool `json:"up"`
	Key     int  `json:"key"`
}

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"username¨`
	Password string    `json:"password"`
	Date     time.Time `json:"date"`
	Email    string    `json:"email"`
	// One to Many relationshionship
	/// USE gorm later
	//Comments []Comment `json:"comments"`
}

func MockUsers(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO users(user_name, password, create_date, email)" +
		"VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec("Simone", "lulz",
		time.Now(), "simone@what.com")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec("AnonBOTTTT", "password",
		time.Now(), "anon@what.com")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func MockVote(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO votes(comment_id, user_id, up)VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	// True means is upvote
	_, err = stmt.Exec(2, 2, true)
	if err != nil {
		fmt.Println(err)
	}

	// True means is upvote
	_, err = stmt.Exec(1, 1, true)
	if err != nil {
		fmt.Println(err)
	}
	// True means is upvote
	_, err = stmt.Exec(1, 2, false)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
