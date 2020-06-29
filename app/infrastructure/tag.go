package infrastructure

import (
	"app/helpers"
	"app/logs"
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

func (t *Tag) CreateTag(tag *models.Tag) (*models.Tag, error) {
	const query = `
INSERT INTO tags (tag) VALUES (?);
`

	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(tag.Tag)
	err = helpers.CheckDBExecError(result, err)
	if err != nil {
		return nil, err
	}

	tagID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	tag.ID = int(tagID)

	logs.Info("Tag created, ID: %v", int(tagID))

	return tag, nil
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

	logs.Info("Tag got, id: %v", tagID)

	return tag, nil
}

func (t *Tag) AttachTagToPin(tagID int, pinID int) error {
	const query = `
INSERT INTO pins_tags (pin_id, tag_id) VALUES (?, ?);
`

	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(pinID, tagID)
	err = helpers.CheckDBExecError(result, err)
	if err != nil {
		return err
	}

	logs.Info("Tag attached to pin, tagID: %v, pinID: %v", tagID, pinID)

	return nil
}

func (t *Tag) GetTagsByPinID(pinID int) ([]*models.Tag, error) {
	const query = `
SELECT t.id, t.tag, t.created_at, t.updated_at FROM tags t JOIN pins_tags pt JOIN pins p WHERE t.id = pt.tag_id AND pt.pin_id = p.id AND p.id = ?;
`

	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(pinID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]*models.Tag, 0)
	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(
			&tag.ID,
			&tag.Tag,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	logs.Info("Tags got, pinID: %v", pinID)

	return tags, nil
}
