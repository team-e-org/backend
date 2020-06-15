package authz

import (
	"app/db"
	"app/mocks"
	"app/models"
	"testing"
)

func TestAuthLayer_AuthenticateUser(t *testing.T) {
	tests := []struct {
		registeredEmail string
		hashedPassword  string
		email           string
		password        string
		wantError       bool
	}{
		{
			"abc@example.com",
			"$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
			"abc@example.com",
			"password",
			false,
		},
		{
			"abc@example.com",
			"$2a$10$a75046RPLEzoN4ObqiOS7Oh/NcvWagI68GHvygZpj3DMvCAekU/1W",
			"abc@example.com",
			"password",
			true,
		},
		{
			"abc@example.com",
			"$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
			"123@example.com",
			"password",
			true,
		},
	}

	for _, tt := range tests {
		user := &models.User{
			ID:             0,
			Email:          tt.registeredEmail,
			HashedPassword: tt.hashedPassword,
		}
		userRepo := mocks.NewUserRepository()
		_ = userRepo.CreateUser(user)
		storage := &db.DataStorage{
			Users: userRepo,
		}
		al := NewAuthLayer(*storage)

		token, err := al.AuthenticateUser(tt.email, tt.password)
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
