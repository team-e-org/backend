package mocks

import (
	"app/models"
	"reflect"
	"testing"
	"time"
)

func TestUserMock(t *testing.T) {
	ID := 0
	users := &UserMock{}
	user := &models.User{
		ID:        ID,
		Name:      "test user",
		Email:     "test@test.com",
		Icon:      "testicon",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users.AddUser(user)
	got, err := users.GetUser(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*user, *got) {
		t.Fatalf("Not equal user")
	}
}

func TestUserMockRepository(t *testing.T) {
	users := NewUserRepository()
	ID := 0
	user := &models.User{
		ID:        ID,
		Name:      "test user",
		Email:     "test@test.com",
		Icon:      "testicon",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	users.AddUser(user)
	got, err := users.GetUser(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*user, *got) {
		t.Fatalf("Not equal user")
	}
}
