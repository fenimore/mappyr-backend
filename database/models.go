package database

import (
	"fmt"
	"time"
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

type User struct {
	Id   int       `json:"id"`
	Name string    `json:"username¨`
	Date time.Time `json:"date"`
	// One to Many relationshionship
	/// USE gorm later
	//Comments []Comment `json:"comments"`
}
