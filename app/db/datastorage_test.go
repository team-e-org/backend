package db

import (
	"app/infrastructure"
	"app/mocks"
	"app/models"
	"app/ptr"
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
	repository.Users().CreateUser(user)
	gotUser, err := repository.Users().GetUser(userID)
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
		Description: ptr.NewString("test description"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repository.Boards().CreateBoard(board)
	gotBoard, err := repository.Boards().GetBoard(boardID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*board, *gotBoard) {
		t.Fatalf("Not equal board")
	}

	pinID := 0
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
	repository.Pins().CreatePin(pin, boardID)
	gotPin, err := repository.Pins().GetPin(pinID)
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
	repository.Tags().CreateTag(tag)
	gotTag, err := repository.Tags().GetTag(tagID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*tag, *gotTag) {
		t.Fatalf("Not equal tag")
	}
}

func TestNewDataStorage(t *testing.T) {
	sqlDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	s3 := infrastructure.NewAWSS3Mock()

	data := NewDataStorage(sqlDB, s3, nil)
	if data.DB() == nil {
		t.Fatalf("DB is nil")
	}
	if data.Users() == nil {
		t.Fatalf("nil field, Users")
	}
	if data.Boards() == nil {
		t.Fatalf("nil field, Boards")
	}
	if data.Pins() == nil {
		t.Fatalf("nil field, Pins")
	}
	if data.Tags() == nil {
		t.Fatalf("nil field, Tags")
	}
	if data.BoardsPins() == nil {
		t.Fatalf("nil field, BoardsPins")
	}
	if data.AWSS3() == nil {
		t.Fatalf("nil field, AWSS3")
	}
}

func BenchmarkDataStorage(b *testing.B) {
	var dataStorage *DataStorage
	var dataStorageInterface DataStorageInterface
	dataStorage = &DataStorage{
		users: mocks.NewUserRepository(),
	}
	dataStorageInterface = &DataStorage{
		users: mocks.NewUserRepository(),
	}

	userID := 1
	user := &models.User{
		ID:             userID,
		Name:           "test user",
		Email:          "test@test.com",
		Icon:           "test icon",
		HashedPassword: "password",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	b.Run("Method call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dataStorage.users.CreateUser(user)
			dataStorage.users.GetUser(userID)
			dataStorage.users.DeleteUser(userID)
		}
	})

	b.Run("Method call via interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dataStorageInterface.Users().CreateUser(user)
			dataStorageInterface.Users().GetUser(userID)
			dataStorageInterface.Users().DeleteUser(userID)
		}
	})
}
