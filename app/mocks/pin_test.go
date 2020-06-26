package mocks

import (
	"app/models"
	"app/ptr"
	"reflect"
	"testing"
	"time"
)

func TestPinMock(t *testing.T) {
	ID := 0
	UserID := 0
	boardID := 0
	pins := &PinMock{}
	pins.ExpectedPins = make([]*models.Pin, 0)
	pins.BoardPinMapper = make(map[int][]int)
	pin := &models.Pin{
		ID:          ID,
		UserID:      ptr.NewInt(UserID),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	pins.CreatePin(pin, boardID)
	got, err := pins.GetPin(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*pin, *got) {
		t.Fatalf("Not equal pin")
	}
}

func TestPinMockRepository(t *testing.T) {
	pins := NewPinRepository()
	ID := 0
	UserID := 0
	boardID := 0
	pin := &models.Pin{
		ID:          ID,
		UserID:      ptr.NewInt(UserID),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	pins.CreatePin(pin, boardID)
	got, err := pins.GetPin(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*pin, *got) {
		t.Fatalf("Not equal pin")
	}
}

func TestPin(t *testing.T) {
	pins := NewPinRepository()
	pin := testBuildPin(t, 0, 0)
	pins.CreatePin(pin, 0)
	p, err := pins.GetPin(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if !reflect.DeepEqual(pin, p) {
		t.Fatalf("Pin did not insert")
	}
	pin2 := testBuildPin(t, 0, 1)
	pins.UpdatePin(pin2)
	p, err = pins.GetPin(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if reflect.DeepEqual(p, pin) || !reflect.DeepEqual(p, pin2) {
		t.Fatalf("Pin did not update")
	}

	pin3 := testBuildPin(t, 1, 1)
	pins.CreatePin(pin3, 0)
	pin4 := testBuildPin(t, 2, 1)
	pins.CreatePin(pin4, 1)

	ps, err := pins.GetPinsByBoardID(0, 0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if ps == nil {
		t.Fatalf("Pins are nil error")
	}
	if len(ps) != 2 {
		t.Fatalf("len(ps) should be 2")
	}
	if !reflect.DeepEqual(pin2, ps[0]) {
		t.Fatalf("pins do not match error")
	}
	if !reflect.DeepEqual(pin3, ps[1]) {
		t.Fatalf("pins do not match error")
	}

	ps, err = pins.GetPinsByBoardID(1, 0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if ps == nil {
		t.Fatalf("Pins are nil error")
	}
	if len(ps) != 1 {
		t.Fatalf("len(ps) should be 1")
	}
	if !reflect.DeepEqual(pin4, ps[0]) {
		t.Fatalf("pins do not match error")
	}

	ps, err = pins.GetPinsByUserID(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if len(ps) != 0 {
		t.Fatalf("len(ps) should be 0")
	}

	ps, err = pins.GetPinsByUserID(1)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if len(ps) != 3 {
		t.Fatalf("len(ps) should be 3")
	}
	if !reflect.DeepEqual(pin2, ps[0]) {
		t.Fatalf("pins do not match error")
	}
	if !reflect.DeepEqual(pin3, ps[1]) {
		t.Fatalf("pins do not match error")
	}
	if !reflect.DeepEqual(pin4, ps[2]) {
		t.Fatalf("pins do not match error")
	}

	ps, err = pins.GetPins(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if len(ps) != 3 {
		t.Fatalf("len(ps) should be 3")
	}
	if !reflect.DeepEqual(pin2, ps[0]) {
		t.Fatalf("pins do not match error")
	}
	if !reflect.DeepEqual(pin3, ps[1]) {
		t.Fatalf("pins do not match error")
	}
	if !reflect.DeepEqual(pin4, ps[2]) {
		t.Fatalf("pins do not match error")
	}

	err = pins.DeletePin(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	ps, err = pins.GetPins(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if len(ps) != 2 {
		t.Fatalf("len(ps) should be 2")
	}
	if !reflect.DeepEqual(pin3, ps[0]) {
		t.Fatalf("pins do not match error")
	}
	if !reflect.DeepEqual(pin4, ps[1]) {
		t.Fatalf("pins do not match error")
	}
}

func testBuildPin(t *testing.T, pinID int, userID int) *models.Pin {
	pin := &models.Pin{
		ID:          pinID,
		UserID:      ptr.NewInt(userID),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return pin
}
