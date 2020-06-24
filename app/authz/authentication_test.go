package authz

import (
	"app/db"
	"app/models"
	"reflect"
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
		storage := db.NewRepositoryMock()
		storage.Users().CreateUser(user)
		al := NewAuthLayerMock(storage)

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

func TestAuthLayer_GetTokenData(t *testing.T) {
	test := struct {
		registeredEmail string
		hashedPassword  string
		email           string
		password        string
	}{
		"abc@example.com",
		"$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
		"abc@example.com",
		"password",
	}

	user := &models.User{
		ID:             0,
		Email:          test.registeredEmail,
		HashedPassword: test.hashedPassword,
	}
	storage := db.NewRepositoryMock()
	storage.Users().CreateUser(user)
	al := NewAuthLayerMock(storage)

	token, _ := al.AuthenticateUser(test.email, test.password)

	_, err := al.GetTokenData("")
	if err == nil {
		t.Error("Error should occur")
	}

	_, err = al.GetTokenData("tekitou")
	if err == nil {
		t.Error("Error should occur")
	}

	tokenData, _ := al.GetTokenData(token)
	if !reflect.DeepEqual(tokenData.UserData, user) {
		t.Error("different token data")
	}
}
