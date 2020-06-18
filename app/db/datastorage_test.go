package db

import (
	"app/models"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository(t *testing.T) {
	repository := NewRepositoryMock()

	userID := 0
	user := &models.User{
		ID:        userID,
		Name:      "test user",
		Email:     "test@test.com",
		Icon:      "test icon",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repository.Users.CreateUser(user)
	gotUser, err := repository.Users.GetUser(userID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*user, *gotUser) {
		t.Fatalf("Not equal user")
	}

	boardID := 0
	board := &models.Board{
		ID:          boardID,
		UserID:      userID,
		Name:        "test board",
		Description: "test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repository.Boards.CreateBoard(board)
	gotBoard, err := repository.Boards.GetBoard(boardID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *gotBoard) {
		t.Fatalf("Not equal board")
	}

	pinID := 0
	pin := &models.Pin{
		ID:          pinID,
		UserID:      userID,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repository.Pins.CreatePin(pin, boardID)
	gotPin, err := repository.Pins.GetPin(pinID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*pin, *gotPin) {
		t.Fatalf("Not equal pin")
	}

	tagID := 0
	tag := &models.Tag{
		ID:        tagID,
		Tag:       "test tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repository.Tags.CreateTag(tag)
	gotTag, err := repository.Tags.GetTag(tagID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*tag, *gotTag) {
		t.Fatalf("Not equal tag")
	}
}

func mockDBHandlingWrapper() (*sql.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	return sqlDB, mock
}
