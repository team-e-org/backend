package models

import "time"

type Board struct {
	ID          int
	UserID      int
	Name        string
	Description string
	isPrivate   bool
	isArchive   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
