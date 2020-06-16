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
	boards.CreateBoard(board)
	got, err := boards.GetBoard(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *got) {
		t.Fatalf("Not equal board")
	}
}

func TestBoardMock_GetBoardsByUserID(t *testing.T) {
	ID := 0
	UserID := 0
	mock := &BoardMock{}
	board := &models.Board{
		ID:          ID,
		UserID:      UserID,
		Name:        "test board",
		Description: "test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mock.CreateBoard(board)
	got, err := mock.GetBoardsByUserID(UserID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *got[0]) {
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
	boards.CreateBoard(board)
	got, err := boards.GetBoard(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *got) {
		t.Fatalf("Not equal board")
	}
}

func TestBoard(t *testing.T) {
	boards := NewBoardRepository()
	ID := 0
	UserID := 0
	now := time.Now()
	board := &models.Board{
		ID:          ID,
		UserID:      UserID,
		Name:        "test board",
		Description: "test description",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	boards.CreateBoard(board)
	board2 := &models.Board{
		ID:          ID,
		UserID:      UserID,
		Name:        "test2 board",
		Description: "test2 description",
		CreatedAt:   now,
		UpdatedAt:   time.Now(),
	}
	boards.UpdateBoard(board2)
	b, err := boards.GetBoard(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if testCompareBoards(t, b, board) {
		t.Fatalf("The board did not update")
	}
	if !testCompareBoards(t, b, board2) {

	}
}

func testCompareBoards(t *testing.T, board *models.Board, board2 *models.Board) bool {
	if board.ID != board2.ID {
		return false
	}
	if board.UserID != board2.UserID {
		return false
	}
	if board.Name != board2.Name {
		return false
	}
	if board.Description != board2.Description {
		return false
	}
	return true
}
