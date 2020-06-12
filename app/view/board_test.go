package view

import (
	"app/models"
	"testing"
	"time"
)

func TestBoard(t *testing.T) {
	b := &models.Board{
		ID:          0,
		UserID:      0,
		Name:        "test name",
		Description: "test description",
		IsPrivate:   false,
		IsArchive:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	v := NewBoard(b)
	if b.ID != v.ID {
		t.Fatalf("ID does not match, model: %v, view: %v", b.ID, v.ID)
	}

	if b.UserID != v.UserID {
		t.Fatalf("UserID does not match, model: %v, view: %v", b.UserID, v.UserID)
	}

	if b.Name != v.Name {
		t.Fatalf("Name does not match, model: %v, view: %v", b.Name, v.Name)
	}

	if b.Description != v.Description {
		t.Fatalf("Description does not match, model: %v, view: %v", b.Description, v.Description)
	}

	if b.IsPrivate != v.IsPrivate {
		t.Fatalf("IsPrivate does not match, model: %v, view: %v", b.IsPrivate, v.IsPrivate)
	}

	if b.IsArchive != v.IsArchive {
		t.Fatalf("IsArchive does not match, model: %v, view: %v", b.IsArchive, v.IsArchive)
	}
}
