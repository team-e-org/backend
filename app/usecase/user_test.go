package usecase

import (
	"app/authz"
	"app/db"
	"app/models"
	"testing"
	"time"
)

func TestUserBoard(t *testing.T) {
	data := db.NewRepositoryMock()
	boards := []*models.Board{
		{
			ID:          0,
			UserID:      0,
			Name:        "test board",
			Description: "test description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          1,
			UserID:      1,
			Name:        "test board2",
			Description: "test description2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      0,
			Name:        "test board3",
			Description: "test description3",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	for _, b := range boards {
		_, err := data.Boards.CreateBoard(b)
		if err != nil {
			t.Fatalf("An error occurred: %v", err)
		}
	}
	authz := authz.NewAuthLayer(data)
	userID := 0
	currentUserID := 0
	boards, err := UserBoards(data, authz, userID, currentUserID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	for _, b := range boards {
		if b.UserID != userID {
			t.Fatalf("Could not get user's boards")
		}
	}
}

func TestEmptyUserBoardError(t *testing.T) {
	data := db.NewRepositoryMock()
	authz := authz.NewAuthLayer(data)
	userID := 0
	currentUserID := 0
	_, err := UserBoards(data, authz, userID, currentUserID)
	if err == nil {
		t.Fatalf("Board is empty")
	}
}

func TestRemovePrivateBoards(t *testing.T) {
	boards := []*models.Board{
		{
			ID:          0,
			UserID:      0,
			Name:        "test board",
			Description: "test description",
			IsPrivate:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          1,
			UserID:      0,
			Name:        "test board2",
			Description: "test description2",
			IsPrivate:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      1,
			Name:        "test board3",
			Description: "test description3",
			IsPrivate:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			UserID:      1,
			Name:        "test board4",
			Description: "test description4",
			IsPrivate:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	userID := 0
	boards = removePrivateBoards(boards, userID)

	if len(boards) != 3 {
		t.Fatalf("len(boards) should be 3")
	}

	for _, b := range boards {
		if b.UserID != userID && b.IsPrivate {
			t.Fatalf("Other people's private boards are gotten.")
		}
	}
}
