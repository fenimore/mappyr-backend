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

type Upvote struct {
	Comment int `json:"comment_id"`
	User    int `json:"user_id"`
	Key     int `json:"key"`
}

type Downvote struct {
	Comment int `json:"comment_id"`
	User    int `json:"user_id"`
	Key     int `json:"key"`
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
	_, err = stmt.Exec("Fenimore", "H2cA45x090",
		time.Now(), "me@what.com")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec("Geffry", "password",
		time.Now(), "he@what.com")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func MockUpvote(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO upvotes(comment_id, user_id)VALUES($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(1, 1)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
