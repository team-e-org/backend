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
	users.CreateUser(user)
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
	users.CreateUser(user)
	got, err := users.GetUser(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*user, *got) {
		t.Fatalf("Not equal user")
	}
}

func TestUser(t *testing.T) {
	users := NewUserRepository()
	ID := 0
	Email := "test@test.com"
	now := time.Now()
	user := &models.User{
		ID:        ID,
		Name:      "test user",
		Email:     Email,
		Icon:      "testicon",
		CreatedAt: now,
		UpdatedAt: now,
	}
	users.CreateUser(user)
	Email2 := "test2@test2.com"
	user2 := &models.User{
		ID:        ID,
		Name:      "test2 user",
		Email:     Email2,
		Icon:      "test2icon",
		CreatedAt: now,
		UpdatedAt: time.Now(),
	}
	users.UpdateUser(user2)
	u, err := users.GetUser(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	users.UpdateUser(user2)
	if testCompareUsers(t, user, u) {
		t.Fatalf("User did not update")
	}
	if !testCompareUsers(t, user2, u) {
		t.Fatalf("User did not update")
	}

	u, err = users.GetUserByEmail(Email)
	if err == nil {
		t.Fatalf("Invalid email, but got user")
	}
	u, err = users.GetUserByEmail(Email2)
	if err != nil {
		t.Fatalf("Valid email, but got error")
	}
	if !testCompareUsers(t, user2, u) {
		t.Fatalf("Users don't match")
	}

	users.DeleteUser(ID)
	u, err = users.GetUser(ID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if u != nil {
		t.Fatalf("User did not delete")
	}
}

func testCompareUsers(t *testing.T, user *models.User, user2 *models.User) bool {
	if user.ID != user2.ID {
		return false
	}
	if user.Name != user2.Name {
		return false
	}
	if user.Email != user2.Email {
		return false
	}
	if user.HashedPassword != user2.HashedPassword {
		return false
	}
	if user.Icon != user2.Icon {
		return false
	}
	return true
}

func TestUserError(t *testing.T) {
	user := &models.User{
		ID:    1,
		Name:  "test name",
		Email: "test@test.com",
		Icon:  "test icon",
	}
	users := NewUserRepository()
	err := users.UpdateUser(user)
	if err == nil {
		t.Fatalf("An error should occur")
	}
	err = users.DeleteUser(user.ID)
	if err == nil {
		t.Fatalf("An error should occur")
	}
}
