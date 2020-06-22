package infrastructure

import (
	"app/models"
	"app/ptr"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateBoard(t *testing.T) {
	userID := 0
	board := &models.Board{
		ID:          0,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO boards (user_id, name, description, is_private) VALUES (?, ?, ?, ?)"))
	prepare.ExpectExec().
		WithArgs(board.UserID, board.Name, board.Description, board.IsPrivate).
		WillReturnResult(sqlmock.NewResult(0, 1))

	boards := NewBoardRepository(sqlDB)
	_, err := boards.CreateBoard(board)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestCreateBoardError(t *testing.T) {
	userID := 0
	board := &models.Board{
		ID:          0,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO boards (user_id, name, description, is_private) VALUES (?, ?, ?, ?)"))
	prepare.ExpectExec().
		WithArgs(board.UserID, board.Name, board.Description, board.IsPrivate).
		WillReturnError(fmt.Errorf("some error"))

	boards := NewBoardRepository(sqlDB)
	_, err := boards.CreateBoard(board)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestUpdateBoard(t *testing.T) {
	userID := 0
	board := &models.Board{
		ID:          0,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test descriptions"),
		IsPrivate:   false,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("UPDATE boards SET name = ?, description = ?, is_private = ?"))
	prepare.ExpectExec().
		WithArgs(board.Name, board.Description, board.IsPrivate).
		WillReturnResult(sqlmock.NewResult(0, 1))

	boards := NewBoardRepository(sqlDB)
	err := boards.UpdateBoard(board)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestUpdateBoardError(t *testing.T) {
	userID := 0
	board := &models.Board{
		ID:          0,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test descriptions"),
		IsPrivate:   false,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("UPDATE boards SET name = ?, description = ?, is_private = ?"))
	prepare.ExpectExec().
		WithArgs(board.Name, board.Description, board.IsPrivate).
		WillReturnError(fmt.Errorf("some error"))

	boards := NewBoardRepository(sqlDB)
	err := boards.UpdateBoard(board)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestDeleteBoard(t *testing.T) {
	boardID := 0

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM boards WHERE id = ?"))
	prepare.ExpectExec().
		WithArgs(boardID).
		WillReturnResult(sqlmock.NewResult(int64(boardID), 1))

	boards := NewBoardRepository(sqlDB)
	err := boards.DeleteBoard(boardID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestDeleteBoardError(t *testing.T) {
	boardID := 0

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM boards WHERE id = ?"))
	prepare.ExpectExec().
		WithArgs(boardID).
		WillReturnError(fmt.Errorf("some error"))

	boards := NewBoardRepository(sqlDB)
	err := boards.DeleteBoard(boardID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetBoard(t *testing.T) {
	boardID := 0

	prepare := mock.ExpectPrepare("SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at FROM boards b WHERE b.id = ?")
	prepare.ExpectQuery().
		WithArgs(boardID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "description", "is_private", "is_archive", "created_at", "updated_at"}).
			AddRow(boardID, 0, "test board", "test description", false, false, now, now))

	boards := NewBoardRepository(sqlDB)
	_, err := boards.GetBoard(boardID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetBoardError(t *testing.T) {
	boardID := 0

	prepare := mock.ExpectPrepare("SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at FROM boards b WHERE b.id = ?")
	prepare.ExpectQuery().
		WithArgs(boardID).
		WillReturnError(fmt.Errorf("some error"))

	boards := NewBoardRepository(sqlDB)
	_, err := boards.GetBoard(boardID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetBoardsByUserID(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare("SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at FROM boards b WHERE b.user_id = ?")
	prepare.ExpectQuery().
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "description", "is_private", "is_archive", "created_at", "updated_at"}).
			AddRow(0, userID, "test board", "test description", false, false, now, now).
			AddRow(1, userID, "test board2", "test description2", false, false, now, now))

	boards := NewBoardRepository(sqlDB)
	_, err := boards.GetBoardsByUserID(userID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetBoardsByUserIDError(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare("SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at FROM boards b WHERE b.user_id = ?")
	prepare.ExpectQuery().
		WithArgs(userID).
		WillReturnError(fmt.Errorf("some error"))

	boards := NewBoardRepository(sqlDB)
	_, err := boards.GetBoardsByUserID(userID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}
