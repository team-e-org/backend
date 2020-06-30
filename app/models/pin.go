package models

import "time"

type Pin struct {
	ID          int     `dynamo:"pin_id"`
	UserID      *int    `dynamo:"post_user_id"`
	Title       string  `dynamo:"title,omitempty"`
	Description *string `dynamo:"description,omitempty"`
	URL         *string `dynamo:"url"`
	ImageURL    string  `dynamo:"image_url"`
	IsPrivate   bool    `dynamo:"is_private"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
