package view

import (
	"app/models"
	"testing"
	"time"
)

func TestPin(t *testing.T) {
	p := &models.Pin{
		ID:          0,
		UserID:      0,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		IsPrivate:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	v := NewPin(p)
	if p.ID != v.ID {
		t.Fatalf("ID does not match, model: %v, view: %v", p.ID, v.ID)
	}

	if p.UserID != v.UserID {
		t.Fatalf("UserID does not match, model: %v, view: %v", p.UserID, v.UserID)
	}

	if p.Title != v.Title {
		t.Fatalf("Title does not match, model: %v, view: %v", p.Title, v.Title)
	}

	if p.Description != v.Description {
		t.Fatalf("Description does not match, model: %v, view: %v", p.Description, v.Description)
	}

	if p.URL != v.URL {
		t.Fatalf("URL does not match, model: %v, view: %v", p.URL, v.URL)
	}

	if p.ImageURL != v.ImageURL {
		t.Fatalf("ImageURL does not match, model: %v, view: %v", p.ImageURL, v.ImageURL)
	}

	if p.IsPrivate != v.IsPrivate {
		t.Fatalf("IsPrivate does not match, model: %v, view: %v", p.IsPrivate, v.IsPrivate)
	}
}
