package view

import (
	"app/models"
	"app/ptr"
	"testing"
	"time"
)

func TestPin(t *testing.T) {
	p := &models.Pin{
		ID:          0,
		UserID:      ptr.NewInt(0),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		IsPrivate:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	v := NewPin(p)
	if p.ID != v.ID {
		t.Fatalf("ID does not match, model: %v, view: %v", p.ID, v.ID)
	}

	if *p.UserID != v.UserID {
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

func TestNewPins(t *testing.T) {
	pins := []*models.Pin{{
		ID:          0,
		UserID:      ptr.NewInt(0),
		Title:       "test title 1",
		Description: ptr.NewString("test description 1"),
		URL:         ptr.NewString("test url 1"),
		ImageURL:    "test image url 1",
		IsPrivate:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, {
		ID:          0,
		UserID:      ptr.NewInt(0),
		Title:       "test title 2",
		Description: ptr.NewString("test description 2"),
		URL:         ptr.NewString("test url 2"),
		ImageURL:    "test image url 2",
		IsPrivate:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}}

	newPins := NewPins(pins)

	for i, newPin := range newPins {
		p := pins[i]

		if p.ID != newPin.ID {
			t.Fatalf("ID does not match, model: %v, view: %v", p.ID, newPin.ID)
		}

		if *p.UserID != newPin.UserID {
			t.Fatalf("UserID does not match, model: %v, view: %v", p.UserID, newPin.UserID)
		}

		if p.Title != newPin.Title {
			t.Fatalf("Title does not match, model: %v, view: %v", p.Title, newPin.Title)
		}

		if p.Description != newPin.Description {
			t.Fatalf("Description does not match, model: %v, view: %v", p.Description, newPin.Description)
		}

		if p.URL != newPin.URL {
			t.Fatalf("URL does not match, model: %v, view: %v", p.URL, newPin.URL)
		}

		if p.ImageURL != newPin.ImageURL {
			t.Fatalf("ImageURL does not match, model: %v, view: %v", p.ImageURL, newPin.ImageURL)
		}

		if p.IsPrivate != newPin.IsPrivate {
			t.Fatalf("IsPrivate does not match, model: %v, view: %v", p.IsPrivate, newPin.IsPrivate)
		}
	}
}
