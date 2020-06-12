package authz

import (
	"app/db"
	"app/mocks"
	"app/models"
	"testing"
)

func TestAuthLayer_AuthenticateUser(t *testing.T) {
	user := &models.User{
		ID:             0,
		Name:           "naoto",
		HashedPassword: "$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
	}
	userRepo := mocks.NewUserRepository()
	_ = userRepo.CreateUser(user)
	storage := &db.DataStorage{
		Users: userRepo,
	}
	al := NewAuthLayer(*storage)

	token, err := al.AuthenticateUser("naoto", "password")
	if err != nil {
		t.Fatalf("An error occured: %v", err)
	}

	if len(token) != 36 {
		t.Fatalf("invalid token: got %s", token)
	}
}
