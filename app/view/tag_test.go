package view

import (
	"app/models"
	"testing"
	"time"
)

func TestTag(t *testing.T) {
	tt := &models.Tag{
		ID:        0,
		Tag:       "test tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	v := NewTag(tt)
	if tt.ID != v.ID {
		t.Fatalf("ID does not match, model: %v, view: %v", tt.ID, v.ID)
	}

	if tt.Tag != v.Tag {
		t.Fatalf("Tag does not match, model: %v, view: %v", tt.Tag, v.Tag)
	}
}
