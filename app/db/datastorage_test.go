package db

import (
	"app/models"
	"reflect"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	repository := NewRepositoryMock()

	user, err := repository.Users.CreateUser("test user", "test@test.com", "testicon", "testpassword")
	if err != nil {
		t.Fatalf("An error occurred while creating user: %v", err)
	}
	gotUser, err := repository.Users.GetUser(user.ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*user, *gotUser) {
		t.Fatalf("Not equal user")
	}

	boardID := 0
	board := &models.Board{
		ID:          boardID,
		UserID:      user.ID,
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
		UserID:      user.ID,
		Title:       "test title",
		Description: "test description",
		URL:         "test url",
		ImageURL:    "test image url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repository.Pins.CreatePin(pin)
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
