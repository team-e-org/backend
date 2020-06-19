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

func TestNewBoards(t *testing.T) {
	boards := []*models.Board{
		{

			ID:          1,
			UserID:      2,
			Name:        "test name 1",
			Description: "test description 1",
			IsPrivate:   false,
			IsArchive:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{

			ID:          3,
			UserID:      4,
			Name:        "test name 2",
			Description: "test description 2",
			IsPrivate:   false,
			IsArchive:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	newBoards := NewBoards(boards)

	for i, newBoard := range newBoards {
		b := boards[i]

		if newBoard.ID != b.ID {
			t.Fatalf("ID does not match, model: %v, view: %v", b.ID, newBoard.ID)
		}
		if newBoard.UserID != b.UserID {
			t.Fatalf("UserID does not match, model: %v, view: %v", b.UserID, newBoard.UserID)
		}
		if newBoard.Name != b.Name {
			t.Fatalf("Name does not match, model: %v, view: %v", b.Name, newBoard.Name)
		}
		if newBoard.Description != b.Description {
			t.Fatalf("Description does not match, model: %v, view: %v", b.Description, newBoard.Description)
		}
		if newBoard.IsPrivate != b.IsPrivate {
			t.Fatalf("isPrivate does not match, model: %v, view: %v", b.IsPrivate, newBoard.IsPrivate)
		}
		if newBoard.IsArchive != b.IsArchive {
			t.Fatalf("isArchive does not match, model: %v, view: %v", b.IsArchive, newBoard.IsArchive)
		}
	}
}
