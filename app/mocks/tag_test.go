package mocks

import (
	"app/models"
	"reflect"
	"testing"
	"time"
)

func TestTagMock(t *testing.T) {
	ID := 0
	tags := &TagMock{}
	tag := &models.Tag{
		ID:        ID,
		Tag:       "test tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tags.AddTag(tag)
	got, err := tags.GetTag(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*tag, *got) {
		t.Fatalf("Not equal tag")
	}
}

func TestTagMockRepository(t *testing.T) {
	tags := NewTagRepository()
	ID := 0
	tag := &models.Tag{
		ID:        ID,
		Tag:       "test tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tags.AddTag(tag)
	got, err := tags.GetTag(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*tag, *got) {
		t.Fatalf("Not equal tag")
	}
}
