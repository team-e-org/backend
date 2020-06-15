package mocks

import (
	"reflect"
	"testing"
)

func TestUserMock(t *testing.T) {
	ID := 0
	users := &UserMock{}
	user, err := users.CreateUser("test user", "test@test.com", "testicon", "testpassword")
	if err != nil {
		t.Fatalf("An error occurred while creating user: %v", err)
	}
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
	user, err := users.CreateUser("test user", "test@test.com", "testicon", "testpassword")
	if err != nil {
		t.Fatalf("An error occurred while creating user: %v", err)
	}
	got, err := users.GetUser(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*user, *got) {
		t.Fatalf("Not equal user")
	}
}
