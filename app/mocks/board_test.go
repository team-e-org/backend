package mocks

import (
	"app/models"
	"reflect"
	"testing"
	"time"
)

func TestBoardMock(t *testing.T) {
	ID := 0
	UserID := 0
	boards := &BoardMock{}
	board := &models.Board{
		ID:          ID,
		UserID:      UserID,
		Name:        "test board",
		Description: "test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	boards.AddBoard(board)
	got, err := boards.GetBoard(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *got) {
		t.Fatalf("Not equal board")
	}
}

func TestBoardMockRepository(t *testing.T) {
	boards := NewBoardRepository()
	ID := 0
	UserID := 0
	board := &models.Board{
		ID:          ID,
		UserID:      UserID,
		Name:        "test board",
		Description: "test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	boards.AddBoard(board)
	got, err := boards.GetBoard(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *got) {
		t.Fatalf("Not equal board")
	}
}
