package authz

import (
	"app/db"
	"app/models"
	"encoding/json"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestGetUserIDByToken(t *testing.T) {
	data := db.NewRepositoryMock()
	al := NewAuthLayerMock(data)
	token := uuid.NewV4().String()
	id, err := GetUserIDByToken(al, token)
	if err == nil {
		t.Fatalf("An error should occur")
	}
	if id != 0 {
		t.Fatalf("ID should be 0")
	}

	hashedPassword, err := db.HashPassword("test password")
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	user := &models.User{
		ID:             1,
		Name:           "test user",
		Email:          "test@test.com",
		Icon:           "test icon",
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	bytes, err := json.Marshal(&TokenData{
		UserData: user,
	})
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	err = al.TokenStorage().SetTokenData(token, string(bytes))
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	d, err := al.TokenStorage().GetTokenData(token)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	var tokenData TokenData
	if err = json.Unmarshal([]byte(d), &tokenData); err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	if user.ID != tokenData.UserData.ID {
		t.Fatalf("ID should same")
	}
	if user.Name != tokenData.UserData.Name {
		t.Fatalf("Name should same")
	}
	if user.Email != tokenData.UserData.Email {
		t.Fatalf("Email should same")
	}
	if user.Icon != tokenData.UserData.Icon {
		t.Fatalf("Icon should same")
	}
	if user.HashedPassword != tokenData.UserData.HashedPassword {
		t.Fatalf("HashedPassword should same")
	}
}

func TestCheckUserPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := db.HashPassword(password)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	if err := checkUserPassword(password, hashedPassword); err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

}
