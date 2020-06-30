package models

import "time"

type Pin struct {
	ID          int       `dynamo:"pin_id"`
	UserID      *int      `dynamo:"post_user_id"`
	Title       string    `dynamo:"title"`
	Description *string   `dynamo:"description"`
	URL         *string   `dynamo:"url"`
	ImageURL    string    `dynamo:"image_url"`
	IsPrivate   bool      `dynamo:"is_private"`
	CreatedAt   time.Time `dynamo:"created_at"`
	UpdatedAt   time.Time `dynamo:"updated_at"`
}
