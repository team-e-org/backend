package db

import "time"

type Pin struct {
	ID         int
	UserID     int
	Title      string
	Descrition string
	URL        string
	ImageURL   string
	IsPrivate  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
