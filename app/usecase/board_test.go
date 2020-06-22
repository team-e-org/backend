package usecase

import (
	"app/db"
	"app/view"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	data := db.NewRepositoryMock()
	id := 0
	userID := 0
	requestBoard := &view.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board",
		Description: "test description",
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err := CreateBoard(data, requestBoard, userID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}
