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
