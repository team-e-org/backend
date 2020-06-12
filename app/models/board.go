package models

import "time"

type Board struct {
	ID          int
	UserID      int
	Name        string
	Description string
	IsPrivate   bool
	IsArchive   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
