package infrastructure

import (
	"app/helpers"
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

func (t *Tag) CreateTag(tag *models.Tag) error {
	const query = `
INSERT INTO tags (tag) VALUES (?);
`

	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(tag.Tag)
	err = helpers.CheckDBExecError(result, err)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tag) GetTag(tagID int) (*models.Tag, error) {
	const query = `
SELECT t.id, t.tag, t.created_at, t.updated_at FROM tags t WHERE t.id = ?;
`

	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(tagID)

	tag := &models.Tag{}
	err = row.Scan(
		&tag.ID,
		&tag.Tag,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *Tag) AttachTagToPin(tagID int, pinID int) error {
	return nil
}

func (t *Tag) GetTagsByPinID(pinID int) ([]*models.Tag, error) {
	return nil, nil
}
