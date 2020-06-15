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
	tags.CreateTag(tag)
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
	tags.CreateTag(tag)
	got, err := tags.GetTag(ID)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if !reflect.DeepEqual(*tag, *got) {
		t.Fatalf("Not equal tag")
	}
}

func TestTag(t *testing.T) {
	tags := NewTagRepository()
	tag := testBuildTag(t, 0)
	tags.CreateTag(tag)
	tt, err := tags.GetTag(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if !reflect.DeepEqual(tag, tt) {
		t.Fatalf("Tags do not match")
	}
	tag2 := testBuildTag(t, 1)
	tags.CreateTag(tag2)
	tag3 := testBuildTag(t, 2)
	tags.CreateTag(tag3)
	tags.AttachTagToPin(0, 0)
	tags.AttachTagToPin(1, 1)
	tags.AttachTagToPin(2, 0)

	ts, err := tags.GetTagsByPinID(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if ts == nil {
		t.Fatalf("Tags are nil error")
	}
	if len(ts) != 2 {
		t.Fatalf("len(ts) should be 2")
	}
	if !reflect.DeepEqual(tag, ts[0]) || !reflect.DeepEqual(tag3, ts[1]) {
		t.Fatalf("Tags do not match error")
	}

	ts, err = tags.GetTagsByPinID(1)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
	if ts == nil {
		t.Fatalf("Tags are nil error")
	}
	if len(ts) != 1 {
		t.Fatalf("len(ts) should be 1")
	}
	if !reflect.DeepEqual(tag2, ts[0]) {
		t.Fatalf("Tags do not match error")
	}
}

func testBuildTag(t *testing.T, tagID int) *models.Tag {
	return &models.Tag{
		ID:        tagID,
		Tag:       "test tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
