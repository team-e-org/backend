package authz

import (
	"app/db"
	"app/mocks"
	"app/models"
	"testing"
)

func TestAuthLayer_AuthenticateUser(t *testing.T) {
	tests := []struct {
		password       string
		hashedPassword string
		wantError      bool
	}{
		{
			"password",
			"$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
			false,
		},
		{
			"password",
			"$2a$10$a75046RPLEzoN4ObqiOS7Oh/NcvWagI68GHvygZpj3DMvCAekU/1W",
			true,
		},
	}

	for _, tt := range tests {
		user := &models.User{
			ID:             0,
			HashedPassword: tt.hashedPassword,
		}
		userRepo := mocks.NewUserRepository()
		_ = userRepo.CreateUser(user)
		storage := &db.DataStorage{
			Users: userRepo,
		}
		al := NewAuthLayer(*storage)

		token, err := al.AuthenticateUser("abc@example.com", tt.password)
		if !tt.wantError && err != nil {
			t.Fatalf("An error occured: %v", err)
		}

		if tt.wantError && err == nil {
			t.Fatalf("Error should occur")
		}

		const wantTokenLength = 36
		if !tt.wantError && len(token) != wantTokenLength {
			t.Fatalf("invalid token length: got %s", token)
		}
	}
}
