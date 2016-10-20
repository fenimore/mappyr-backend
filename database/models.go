package database

import (
	"fmt"
	"time"

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
	return fmt.Sprintf("Comment %d: %s, %d\nAt %f %f",
		c.Id, c.Title, c.Upvotes, c.Lat, c.Lon)
}

type Vote struct {
	Comment int  `json:"comment_id"`
	User    int  `json:"user_id"`
	Up      bool `json:"up"`
}

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"username"`
	Password string    `json:"password"`
	Date     time.Time `json:"date"`
	Email    string    `json:"email"`
	// One to Many relationshionship
	/// USE gorm later
	//Comments []Comment `json:"comments"`
}
