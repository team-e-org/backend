package usecase

import (
	"app/db"
	"app/models"
	"app/ptr"
	"testing"
	"time"
)

func TestCreateBoard(t *testing.T) {
	data := db.NewRepositoryMock()
	id := 0
	userID := 0
	requestBoard := &models.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err := CreateBoard(data, requestBoard)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestUpdateBoard(t *testing.T) {
	var err error
	data := db.NewRepositoryMock()
	id := 0
	userID := 0
	board := &models.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err = CreateBoard(data, board)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	boardUpdated := &models.Board{
		ID:          id,
		UserID:      userID,
		Name:        "test board updated",
		Description: ptr.NewString("test description updated"),
		IsPrivate:   false,
		IsArchive:   false,
	}
	_, err = UpdateBoard(data, boardUpdated)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestSavePin(t *testing.T) {
	data := db.NewRepositoryMock()
	userID := 1
	boardID := 2
	pinID := 3
	board := &models.Board{
		ID:          boardID,
		UserID:      userID,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		IsArchive:   false,
	}

	_, err := CreateBoard(data, board)
	if err != nil {
		t.Fatalf("An error occurred")
	}

	pin := &models.Pin{
		ID:          pinID,
		UserID:      ptr.NewInt(userID),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		IsPrivate:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = CreatePin(data, pin, nil, "", ".png", boardID)
	if err != nil {
		t.Fatalf("An error occurred")
	}

	tests := []struct {
		name    string
		data    db.DataStorageInterface
		boardID int
		pinID   int
		wantErr bool
	}{
		{
			name:    "save pin",
			data:    data,
			boardID: 2,
			pinID:   3,
			wantErr: false,
		},
		{
			name:    "board doesn't exist",
			data:    data,
			boardID: 99,
			pinID:   3,
			wantErr: true,
		},
		{
			name:    "pin doesn't exist",
			data:    data,
			boardID: 2,
			pinID:   99,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SavePin(tt.data, tt.boardID, tt.pinID); (err != nil) != tt.wantErr {
				t.Errorf("SavePin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnsavePin(t *testing.T) {
	createData := func() db.DataStorageInterface {
		data := db.NewRepositoryMock()
		userID := 1
		boardID := 2
		pinID := 3
		board := &models.Board{
			ID:          boardID,
			UserID:      userID,
			Name:        "test board",
			Description: ptr.NewString("test description"),
			IsPrivate:   false,
			IsArchive:   false,
		}

		_, err := CreateBoard(data, board)
		if err != nil {
			t.Fatalf("An error occurred")
		}

		pin := &models.Pin{
			ID:          pinID,
			UserID:      ptr.NewInt(userID),
			Title:       "test title",
			Description: ptr.NewString("test description"),
			URL:         ptr.NewString("test url"),
			ImageURL:    "test image url",
			IsPrivate:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err = CreatePin(data, pin, nil, "", ".png", boardID)
		if err != nil {
			t.Fatalf("An error occurred")
		}

		err = SavePin(data, 2, 3)
		if err != nil {
			t.Fatalf("An error occurred")
		}

		return data
	}

	tests := []struct {
		name    string
		data    db.DataStorageInterface
		boardID int
		pinID   int
		wantErr bool
	}{
		{
			name:    "success",
			data:    createData(),
			boardID: 2,
			pinID:   3,
			wantErr: false,
		},
		{
			name:    "boardID is wrong",
			data:    createData(),
			boardID: 99,
			pinID:   3,
			wantErr: true,
		},
		{
			name:    "pinID is wrong",
			data:    createData(),
			boardID: 2,
			pinID:   99,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnsavePin(tt.data, tt.boardID, tt.pinID); (err != nil) != tt.wantErr {
				t.Errorf("UnsavePin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
