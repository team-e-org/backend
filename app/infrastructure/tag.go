package infrastructure

import (
	"app/models"
	"app/repository"
	"database/sql"
)

type Tag struct {
	DB *sql.DB
}

func NewTagRepository(db *sql.DB) repository.TagRepository {
	return &Tag{
		DB: db,
	}
}

func (u *Tag) CreateTag(tag *models.Tag) error {
	return nil
}

func (u *Tag) GetTag(tagID int) (*models.Tag, error) {
	return nil, nil
}
