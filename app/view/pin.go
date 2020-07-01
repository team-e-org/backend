package view

import (
	"app/models"
	"app/ptr"
	"time"
)

type Pin struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userId"`
	Title       string  `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
	ImageURL    string  `json:"imageUrl"`
	IsPrivate   bool    `json:"isPrivate"`
}

type DynamoPin struct {
	ID          int    `dynamo:"pin_id"`
	UserID      int    `dynamo:"post_user_id"`
	Title       string `dynamo:"title,omitempty"`
	Description string `dynamo:"description,omitempty"`
	URL         string `dynamo:"url,omitempty"`
	ImageURL    string `dynamo:"image_url"`
	IsPrivate   bool   `dynamo:"is_private"`
}

func NewPin(pin *models.Pin) *Pin {
	p := &Pin{
		ID:          pin.ID,
		UserID:      *pin.UserID,
		Title:       pin.Title,
		Description: pin.Description,
		URL:         pin.URL,
		ImageURL:    pin.ImageURL,
		IsPrivate:   pin.IsPrivate,
	}

	return p
}

func NewPins(pins []*models.Pin) []*Pin {
	b := make([]*Pin, 0, len(pins))

	for _, pin := range pins {
		b = append(b, NewPin(pin))
	}

	return b
}

func NewPinModel(pin *Pin) *models.Pin {
	p := &models.Pin{
		ID:          pin.ID,
		UserID:      ptr.NewInt(pin.UserID),
		Title:       pin.Title,
		Description: pin.Description,
		URL:         pin.URL,
		ImageURL:    pin.ImageURL,
		IsPrivate:   pin.IsPrivate,
	}

	return p
}

type AttachTagsLambdaPayloadPin struct {
	ID          int        `json:"id"`
	UserID      int        `json:"userId"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         *string    `json:"url,omitempty"`
	ImageURL    string     `json:"imageUrl"`
	IsPrivate   bool       `json:"isPrivate"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func NewLambdaPin(pin *models.Pin) *AttachTagsLambdaPayloadPin {
	p := &AttachTagsLambdaPayloadPin{
		ID:          pin.ID,
		UserID:      *pin.UserID,
		Title:       pin.Title,
		Description: *pin.Description,
		URL:         pin.URL,
		ImageURL:    pin.ImageURL,
		IsPrivate:   pin.IsPrivate,
		CreatedAt:   &pin.CreatedAt,
		UpdatedAt:   &pin.UpdatedAt,
	}

	return p
}
