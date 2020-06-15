package view

import "app/models"

type Pin struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	BoardID     int    `json:"boardId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
	IsPrivate   bool   `json:"isPrivate"`
}

func NewPin(pin *models.Pin) *Pin {
	p := &Pin{
		ID:          pin.ID,
		UserID:      pin.UserID,
		Title:       pin.Title,
		Description: pin.Description,
		URL:         pin.URL,
		ImageURL:    pin.ImageURL,
		IsPrivate:   pin.IsPrivate,
	}

	return p
}
