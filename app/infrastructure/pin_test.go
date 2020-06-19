package infrastructure

import (
	"app/models"
	"app/ptr"
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
		URL:         ptr.NewString("test url"),
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
	_, err := pins.CreatePin(pin, boardID)
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
		URL:         ptr.NewString("test url"),
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
	_, err := pins.CreatePin(pin, boardID)
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
		URL:         ptr.NewString("test url"),
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
		URL:         ptr.NewString("test url"),
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

func TestDeletePinError(t *testing.T) {
	id := 1

	mock.ExpectBegin()
	prepare := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM pins WHERE id = ?"))
	prepare.ExpectExec().
		WithArgs().
		WillReturnResult(sqlmock.NewResult(int64(id), 1))

	prepare2 := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM boards_pins WHERE pin_id = ?"))
	prepare2.ExpectExec().
		WithArgs().
		WillReturnError(fmt.Errorf("some error"))

	mock.ExpectRollback()

	pins := NewPinRepository(sqlDB)
	err := pins.DeletePin(id)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetPin(t *testing.T) {
	id := 1

	prepare := mock.ExpectPrepare("SELECT p.id, p.user_id, p.title, p.description, p.url, p.image_url, p.is_private, p.created_at, p.updated_at FROM pins p WHERE p.id = ?")
	prepare.ExpectQuery().
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "description", "url", "image_url", "is_private", "created_at", "updated_at"}).
			AddRow(1, 1, "test title", "test_description", "test url", "test image url", false, now, now))

	pins := NewPinRepository(sqlDB)
	_, err := pins.GetPin(id)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetPinError(t *testing.T) {
	id := 1

	prepare := mock.ExpectPrepare("SELECT p.id, p.user_id, p.title, p.description, p.url, p.image_url, p.is_private, p.created_at, p.updated_at")
	prepare.ExpectQuery().
		WithArgs(id).
		WillReturnError(fmt.Errorf("some error"))

	pins := NewPinRepository(sqlDB)
	_, err := pins.GetPin(id)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

// func TestGetPinsByBoardID(t *testing.T) {
// 	boardID := 1
// 	page := 1
//
// 	prepare := mock.ExpectPrepare("SELECT p.id, p.user_id, p.title, p.description, p.url, p.image_url, p.is_private, p.created_at, p.updated_at FROM pins AS p JOIN boards_pins AS bp ON p.id = bp.pin_id WHERE bp.board_id = $1 LIMIT $2 OFFSET $3")
// 	prepare.ExpectQuery().
// 		WithArgs(boardID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "description", "url", "image_url", "is_private", "created_at", "updated_at"}).
// 			AddRow(1, 1, "test title", "test description", "test url", "test image url", false, now, now).
// 			AddRow(2, 2, "test title2", "test description2", "test url2", "test image url2", false, now, now))
//
// 	pins := NewPinRepository(sqlDB)
// 	_, err := pins.GetPinsByBoardID(boardID, page)
// 	if err != nil {
// 		t.Fatalf("An error occurred: %v\n", err)
// 	}
//
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Unfulfilled expectations error: %v\n", err)
// 	}
// }
//
// func TestGetPinsByBoardIDError(t *testing.T) {
// 	boardID := 1
// 	page := 1
//
// 	prepare := mock.ExpectPrepare("SELECT p.id, p.user_id, p.title, p.description, p.url, p.image_url, p.is_private, p.created_at, p.updated_at FROM pins AS p JOIN boards_pins AS bp ON p.id = bp.pin_id WHERE bp.board_id = $1 LIMIT $2 OFFSET $3")
// 	prepare.ExpectQuery().
// 		WithArgs(boardID).
// 		WillReturnError(fmt.Errorf("some error"))
//
// 	pins := NewPinRepository(sqlDB)
// 	_, err := pins.GetPinsByBoardID(boardID, page)
// 	if err == nil {
// 		t.Fatalf("An error occurred: %v\n", err)
// 	}
//
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("Unfulfilled expectations error: %v\n", err)
// 	}
// }

func TestGetPinsByUserID(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare("SELECT p.id, p.user_id, p.title, p.description, p.url, p.image_url, p.is_private, p.created_at, p.updated_at FROM pins p WHERE p.user_id = ?")
	prepare.ExpectQuery().
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "description", "url", "image_url", "is_private", "created_at", "updated_at"}).
			AddRow(1, 1, "test title", "test description", "test url", "test image url", false, now, now).
			AddRow(2, 1, "test title2", "test description2", "test url2", "test image url2", false, now, now))

	pins := NewPinRepository(sqlDB)
	_, err := pins.GetPinsByUserID(userID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetPinsByUserIDError(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare("SELECT p.id, p.user_id, p.title, p.description, p.url, p.image_url, p.is_private, p.created_at, p.updated_at FROM pins p WHERE p.user_id = ?")
	prepare.ExpectQuery().
		WithArgs(userID).
		WillReturnError(fmt.Errorf("some error"))

	pins := NewPinRepository(sqlDB)
	_, err := pins.GetPinsByUserID(userID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}
