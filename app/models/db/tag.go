package db

import "time"

type Tag struct {
	ID        int
	Tag       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
