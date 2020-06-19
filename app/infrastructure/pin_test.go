package infrastructure

import (
	"app/models"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreatePin(t *testing.T) {
	id := 1
	userID := 1
	boardID := 1
	pin := &models.Pin{
		ID:          id,
		UserID:      userID,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		IsPrivate:   false,
	}

	mock.ExpectBegin()
	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO pins (user_id, title, description, url, image_url, is_private) VALUES (?, ?, ?, ?, ?, ?)"))
	prepare.ExpectExec().
		WithArgs(pin.UserID, pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	prepare2 := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO boards_pins (board_id, pin_id) VALUES (?, ?)"))
	prepare2.ExpectExec().
		WithArgs(boardID, id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	pins := NewPinRepository(sqlDB)
	err := pins.CreatePin(pin, boardID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestCreatePinError(t *testing.T) {
	id := 1
	userID := 1
	boardID := 1
	pin := &models.Pin{
		ID:          id,
		UserID:      userID,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		IsPrivate:   false,
	}

	mock.ExpectBegin()
	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO pins (user_id, title, description, url, image_url, is_private) VALUES (?, ?, ?, ?, ?, ?)"))
	prepare.ExpectExec().
		WithArgs(pin.UserID, pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	prepare2 := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO boards_pins (board_id, pin_id) VALUES (?, ?)"))
	prepare2.ExpectExec().
		WithArgs(boardID, id).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	pins := NewPinRepository(sqlDB)
	err := pins.CreatePin(pin, boardID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestUpdatePin(t *testing.T) {
	id := 1
	userID := 1
	pin := &models.Pin{
		ID:          id,
		UserID:      userID,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		IsPrivate:   false,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("UPDATE pins SET title = ?, description = ?, url = ?, image_url = ?, is_private = ?"))
	prepare.ExpectExec().
		WithArgs(pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	pins := NewPinRepository(sqlDB)
	err := pins.UpdatePin(pin)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestUpdatePinError(t *testing.T) {
	id := 1
	userID := 1
	pin := &models.Pin{
		ID:          id,
		UserID:      userID,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		IsPrivate:   false,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("UPDATE pins SET title = ?, description = ?, url = ?, image_url = ?, is_private = ?"))
	prepare.ExpectExec().
		WithArgs(pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate).
		WillReturnError(fmt.Errorf("some error"))

	pins := NewPinRepository(sqlDB)
	err := pins.UpdatePin(pin)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestDeletePin(t *testing.T) {
	id := 1

	mock.ExpectBegin()
	prepare := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM pins WHERE id = ?"))
	prepare.ExpectExec().
		WithArgs().
		WillReturnResult(sqlmock.NewResult(int64(id), 1))

	prepare2 := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM boards_pins WHERE pin_id = ?"))
	prepare2.ExpectExec().
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	pins := NewPinRepository(sqlDB)
	err := pins.DeletePin(id)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}
