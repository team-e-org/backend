// +build integration

package db

import (
	"app/config"
	"app/mocks"
	"app/models"
	"app/ptr"
	"database/sql"
	"log"
	"os"
	"testing"
)

var sqlDB *sql.DB
var s3 S3
var data DataStorageInterface

func dbHandlingWrapper(m *testing.M) int {
	c, err := config.ReadDBConfig()
	if err != nil {
		panic(err)
	}
	sqlDB, err := ConnectToMySql(*c)
	if err != nil {
		log.Panicf("Can not connect to DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	s3 = mocks.NewAWSS3Repository()
	data = NewDataStorage(sqlDB, s3)

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(dbHandlingWrapper(m))
}

func TestUser(t *testing.T) {
	user := testCreateUser(t)
	user, err := data.Users().CreateUser(user)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	user, err = data.Users().GetUserByEmail(user.Email)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	data.Users().DeleteUser(user.ID)
	user = testCreateUser(t)
	_, err = data.Users().CreateUser(user)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	userGot, err := data.Users().GetUserByEmail(user.Email)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if userGot.Name != user.Name {
		t.Fatalf("Users do not match error")
	}
	user2 := testCreateUser2(t)
	user2.ID = userGot.ID
	err = data.Users().UpdateUser(user2)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	user3, err := data.Users().GetUser(user2.ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if user2.ID != user3.ID {
		t.Fatalf("Users do not match error")
	}
	if user2.Email != user3.Email {
		t.Fatalf("Users do not match error")
	}
	err = data.Users().DeleteUser(user3.ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func testCreateUser(t *testing.T) *models.User {
	return &models.User{
		Name:  "test user",
		Email: "test@test.com",
		Icon:  "test icon",
	}
}

func testCreateUser2(t *testing.T) *models.User {
	return &models.User{
		Name:  "test user2",
		Email: "test2@test.com",
		Icon:  "test icon2",
	}
}

func TestBoard(t *testing.T) {
	board := testCreateBoard(t)
	board, err := data.Boards().CreateBoard(board)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	boardGot, err := data.Boards().GetBoard(board.ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if board.Name != boardGot.Name {
		t.Fatalf("Boards do not match error")
	}
	board2 := testCreateBoard2(t)
	board2.ID = board.ID
	err = data.Boards().UpdateBoard(board2)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	board3, err := data.Boards().GetBoardsByUserID(100)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if board2.Name != board3[0].Name {
		t.Fatalf("Boards do not match error")
	}
	err = data.Boards().DeleteBoard(board3[0].ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func testCreateBoard(t *testing.T) *models.Board {
	return &models.Board{
		UserID:      100,
		Name:        "test board",
		Description: ptr.NewString("test description"),
		IsPrivate:   false,
		IsArchive:   false,
	}
}

func testCreateBoard2(t *testing.T) *models.Board {
	return &models.Board{
		UserID:      100,
		Name:        "test board2",
		Description: ptr.NewString("test description2"),
		IsPrivate:   false,
		IsArchive:   false,
	}
}

func TestPin(t *testing.T) {
	boardID := 100
	pin := testCreatePin(t)
	pin, err := data.Pins().CreatePin(pin, boardID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	pinGot, err := data.Pins().GetPin(pin.ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if pin.Title != pinGot.Title {
		t.Fatalf("Pins do not match error")
	}
	pin2 := testCreatePin2(t)
	pin2.ID = pin.ID
	err = data.Pins().UpdatePin(pin2)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	pin3, err := data.Pins().GetPinsByUserID(100)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if pin2.Title != pin3[0].Title {
		t.Fatalf("Pins do not match error")
	}
	err = data.Pins().DeletePin(pin3[0].ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func testCreatePin(t *testing.T) *models.Pin {
	return &models.Pin{
		UserID:      ptr.NewInt(100),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		IsPrivate:   false,
	}
}

func testCreatePin2(t *testing.T) *models.Pin {
	return &models.Pin{
		UserID:      ptr.NewInt(100),
		Title:       "test title2",
		Description: ptr.NewString("test description2"),
		URL:         ptr.NewString("test url2"),
		ImageURL:    "test image url2",
		IsPrivate:   false,
	}
}

func TestTag(t *testing.T) {
	tag := testCreateTag(t)
	tag, err := data.Tags().CreateTag(tag)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	tagGot, err := data.Tags().GetTag(tag.ID)
	if tag.Tag != tagGot.Tag {
		t.Fatalf("Tags do not match error")
	}
	_, err = data.DB().Exec("DELETE FROM tags WHERE id = ?;", tag.ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func testCreateTag(t *testing.T) *models.Tag {
	return &models.Tag{
		Tag: "test",
	}
}
