package infrastructure

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateBoardPin(t *testing.T) {
	boardID := 0
	pinID := 0

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO boards_pins(board_id, pin_id) VALUES (?, ?)"))
	prepare.
		ExpectExec().
		WithArgs(boardID, pinID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	boardsPins := NewBoardPinRepository(sqlDB)
	err := boardsPins.CreateBoardPin(boardID, pinID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestCreateBoardPinError(t *testing.T) {
	boardID := 0
	pinID := 0

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO boards_pins(board_id, pin_id) VALUES (?, ?)"))
	prepare.
		ExpectExec().
		WithArgs(boardID, pinID).
		WillReturnError(fmt.Errorf("some error"))

	boardsPins := NewBoardPinRepository(sqlDB)
	err := boardsPins.CreateBoardPin(boardID, pinID)
	if err == nil {
		t.Fatalf("Error should occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestBoardPin_DeleteBoardPin(t *testing.T) {
	boardID := 0
	pinID := 0

	query := "DELETE FROM boards_pins WHERE board_id = ? and pin_id = ?"
	prepare := mock.ExpectPrepare(regexp.QuoteMeta(query))

	prepare.ExpectExec().
		WithArgs(boardID, pinID).
		WillReturnResult(sqlmock.NewResult(int64(0), 1))
	boardsPins := NewBoardPinRepository(sqlDB)
	err := boardsPins.DeleteBoardPin(boardID, pinID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v", err)
	}
}
