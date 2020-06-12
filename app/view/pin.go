package view

import "app/models"

type Pin struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	IsPrivate   bool   `json:"is_private"`
}

func NewPin(pin *models.Pin) *Pin {
	p := &Pin{
		ID:          pin.ID,
		UserID:      pin.UserID,
		Title:       pin.Title,
		Description: pin.Descrition,
		URL:         pin.URL,
		ImageURL:    pin.ImageURL,
		IsPrivate:   pin.IsPrivate,
	}

	return p
}
