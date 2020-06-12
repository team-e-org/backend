package models

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
