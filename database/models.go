package database

import (
	"fmt"
	"time"
)

// TODO: Add a place model?

type Comment struct {
	Id     int64     `json:"id"`
	Text   string    `json:"title"`
	Lat    float64   `json:"latitude"`
	Lon    float64   `json:"longitude"`
	Date   time.Time `json:"date"`
	UserId int64     `json:"user-id"` // 0 for anon
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment %i: %s\nAt %f %f",
		s.Id, s.Text, s.Lat, s.Lon)
}

type User struct {
	Id   int64     `json:"id"`
	Name string    `json:"username¨`
	Date time.Time `json:"date"`
	// One to Many relationshionship
	/// USE gorm later
	//Comments []Comment `json:"comments"`
}
