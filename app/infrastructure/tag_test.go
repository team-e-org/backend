package infrastructure

import (
	"app/models"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTag(t *testing.T) {
	id := 1
	tag := &models.Tag{
		ID:  id,
		Tag: "test tag",
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO tags (tag) VALUES (?)"))
	prepare.ExpectExec().
		WithArgs(tag.Tag).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tags := NewTagRepository(sqlDB)
	_, err := tags.CreateTag(tag)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestCreateTagError(t *testing.T) {
	id := 1
	tag := &models.Tag{
		ID:  id,
		Tag: "test tag",
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO tags (tag) VALUES (?)"))
	prepare.ExpectExec().
		WithArgs(tag.Tag).
		WillReturnError(fmt.Errorf("some error"))

	tags := NewTagRepository(sqlDB)
	_, err := tags.CreateTag(tag)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetTag(t *testing.T) {
	id := 1

	prepare := mock.ExpectPrepare("SELECT t.id, t.tag, t.created_at, t.updated_at FROM tags t WHERE t.id = ?")
	prepare.ExpectQuery().
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "tag", "created_at", "updated_at"}).
			AddRow(id, "test tag", now, now))

	tags := NewTagRepository(sqlDB)
	_, err := tags.GetTag(id)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetTagError(t *testing.T) {
	id := 1

	prepare := mock.ExpectPrepare("SELECT t.id, t.tag, t.created_at, t.updated_at FROM tags t WHERE t.id = ?")
	prepare.ExpectQuery().
		WithArgs(id).
		WillReturnError(fmt.Errorf("some error"))

	tags := NewTagRepository(sqlDB)
	_, err := tags.GetTag(id)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestAttachTagToPin(t *testing.T) {
	id := 1
	pinID := 1

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO pins_tags (pin_id, tag_id) VALUES (?, ?)"))
	prepare.ExpectExec().
		WithArgs(pinID, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tags := NewTagRepository(sqlDB)
	err := tags.AttachTagToPin(id, pinID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestAttachTagToPinError(t *testing.T) {
	id := 1
	pinID := 1

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO pins_tags (pin_id, tag_id) VALUES (?, ?)"))
	prepare.ExpectExec().
		WithArgs(pinID, id).
		WillReturnError(fmt.Errorf("some error"))

	tags := NewTagRepository(sqlDB)
	err := tags.AttachTagToPin(id, pinID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetTagsByPinID(t *testing.T) {
	pinID := 1

	prepare := mock.ExpectPrepare("SELECT t.id, t.tag, t.created_at, t.updated_at FROM tags t JOIN pins_tags pt JOIN pins p WHERE t.id = pt.tag_id AND pt.pin_id = p.id AND p.id = ?")
	prepare.ExpectQuery().
		WithArgs(pinID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "tag", "created_at", "updated_at"}).
			AddRow(1, "test tag", now, now).
			AddRow(2, "test tag2", now, now))

	tags := NewTagRepository(sqlDB)
	_, err := tags.GetTagsByPinID(pinID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetTagsByPinIDError(t *testing.T) {
	pinID := 1

	prepare := mock.ExpectPrepare("SELECT t.id, t.tag, t.created_at, t.updated_at FROM tags t JOIN pins_tags pt JOIN pins p WHERE t.id = pt.tag_id AND pt.pin_id = p.id AND p.id = ?")
	prepare.ExpectQuery().
		WithArgs(pinID).
		WillReturnError(fmt.Errorf("some error"))

	tags := NewTagRepository(sqlDB)
	_, err := tags.GetTagsByPinID(pinID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}
