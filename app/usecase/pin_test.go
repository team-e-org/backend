package usecase

import (
	"app/db"
	"app/models"
	"app/ptr"
	"testing"
	"time"
)

func createPins(t *testing.T) []*models.Pin {
	pins := []*models.Pin{
		{
			ID:          0,
			UserID:      ptr.NewInt(0),
			Title:       "test title",
			Description: ptr.NewString("test description"),
			URL:         ptr.NewString("test url"),
			ImageURL:    "test image url",
			IsPrivate:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          1,
			UserID:      ptr.NewInt(0),
			Title:       "test title2",
			Description: ptr.NewString("test description2"),
			URL:         ptr.NewString("test url2"),
			ImageURL:    "test image url2",
			IsPrivate:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      ptr.NewInt(1),
			Title:       "test title3",
			Description: ptr.NewString("test description3"),
			URL:         ptr.NewString("test url3"),
			ImageURL:    "test image url3",
			IsPrivate:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          3,
			UserID:      ptr.NewInt(1),
			Title:       "test title4",
			Description: ptr.NewString("test description4"),
			URL:         ptr.NewString("test url4"),
			ImageURL:    "test image url4",
			IsPrivate:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return pins
}

func insertPins(t *testing.T, data db.DataStorageInterface, pins []*models.Pin, boardID int) {
	for i, p := range pins {
		pp, err := CreatePin(data, p, boardID)
		if err != nil {
			t.Fatalf("An error occurred: %v", err)
		}
		if pp != pins[i] {
			t.Fatalf("A pin not created")
		}
	}
}

func TestCreatePin(t *testing.T) {
	boardID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, boardID)
}

func TestUpdatePin(t *testing.T) {
	boardID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, boardID)
	newPin := &models.Pin{
		ID:          0,
		Title:       "test title updated",
		Description: ptr.NewString("test description updated"),
		URL:         ptr.NewString("test url updated"),
	}
	err := data.Pins().UpdatePin(newPin)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	pin, err := data.Pins().GetPin(0)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if pin.ID != newPin.ID {
		t.Fatalf("PinID does not match error")
	}
	if pin.Title != newPin.Title {
		t.Fatalf("Pin title does not match error")
	}
	if *pin.Description != *newPin.Description {
		t.Fatalf("Pin description does not match error")
	}
	if *pin.URL != *newPin.URL {
		t.Fatalf("Pin URL does not match error")
	}
}

func TestRemovePrivatePins(t *testing.T) {
	userID := 0
	pins := createPins(t)
	pins = removePrivatePin(pins, userID)
	if len(pins) != 3 {
		t.Fatalf("len(pins) should be 3")
	}
	for _, p := range pins {
		if p.IsPrivate && *p.UserID != userID {
			t.Fatalf("Other people's private pins are gotten")
		}
	}
}

func TestServePin(t *testing.T) {
	userID := 0
	boardID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, boardID)
	_, err := ServePin(data, 0, userID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	_, err = ServePin(data, 1, userID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	_, err = ServePin(data, 2, userID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestServePinError(t *testing.T) {
	userID := 0
	boardID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, boardID)
	_, err := ServePin(data, 3, userID)
	if err == nil {
		t.Fatalf("Got other people's private pin")
	}
	_, err = ServePin(data, 4, userID)
	if err == nil {
		t.Fatalf("Pin should be nil")
	}
}

func TestGetPins(t *testing.T) {
	boardID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, boardID)
	pins, err := GetPins(data, 0)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if len(pins) != 2 {
		t.Fatalf("len(pins) should be 2")
	}
}

func TestGetPinsByBoardID(t *testing.T) {
	userID := 0
	boardID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, boardID)
	pins, err := GetPinsByBoardID(data, userID, boardID, 0)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if len(pins) != 3 {
		t.Fatalf("len(pins) should be 3")
	}
}

func TestGetPinsByBoardIDError(t *testing.T) {
	userID := 0
	data := db.NewRepositoryMock()
	pins := createPins(t)
	insertPins(t, data, pins, 0)
	_, err := GetPinsByBoardID(data, userID, 1, 0)
	if err == nil {
		t.Fatalf("Board shoud not exist")
	}
}
