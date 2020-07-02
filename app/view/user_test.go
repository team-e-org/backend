package view

import (
	"app/models"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	u := &models.User{
		ID:             0,
		Name:           "test name",
		Email:          "test email",
		Icon:           "test icon",
		HashedPassword: "test password",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	v := NewUser(u)
	if u.ID != v.ID {
		t.Fatalf("ID does not match, model: %v, view: %v", u.ID, v.ID)
	}

	if u.Name != v.Name {
		t.Fatalf("Name does not match, model: %v, view: %v", u.Name, v.Name)
	}

	if u.Email != v.Email {
		t.Fatalf("Email does not match, model: %v, view: %v", u.Email, v.Email)
	}

	if u.Icon != v.Icon {
		t.Fatalf("Icon does not match, model: %v, view: %v", u.Icon, v.Icon)
	}
}

func TestNewUserModel(t *testing.T) {
	vu := &User{
		ID:       0,
		Name:     "test name",
		Email:    "test email",
		Icon:     "test icon",
		Password: "test password",
	}

	mu, err := NewUserModel(vu)
	if err != nil {
		t.Fatalf("Error occured: %v", err)
	}

	if vu.ID != mu.ID {
		t.Fatalf("ID does not match, view: %v, model: %v", mu.ID, vu.ID)
	}

	if vu.Name != mu.Name {
		t.Fatalf("Name does not match, view: %v, model: %v", vu.Name, mu.Name)
	}

	if vu.Email != mu.Email {
		t.Fatalf("Email does not match, view: %v, model: %v", vu.Email, mu.Email)
	}

	if vu.Icon != mu.Icon {
		t.Fatalf("Icon does not match, view: %v, model: %v", vu.Icon, mu.Icon)
	}
}
