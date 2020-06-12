package models

import "time"

type Pin struct {
	ID          int
	UserID      int
	Title       string
	Description string
	URL         string
	ImageURL    string
	IsPrivate   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
