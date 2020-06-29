package usecase

import (
	"app/db"
	"app/models"
	"app/ptr"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	data := db.NewRepositoryMock()
	id := 0
	userID := 0
	requestBoard := &models.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err := CreateBoard(data, requestBoard)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestUpdateBoard(t *testing.T) {
	var err error
	data := db.NewRepositoryMock()
	id := 0
	userID := 0
	board := &models.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err = CreateBoard(data, board)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	boardUpdated := &models.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board updated",
		Description: ptr.NewString("test description updated"),
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err = UpdateBoard(data, boardUpdated)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}
