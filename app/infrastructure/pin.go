package infrastructure

import (
	"app/helpers"
	"app/models"
	"app/repository"
	"database/sql"
)

type Pin struct {
	DB *sql.DB
}

func NewPinRepository(db *sql.DB) repository.PinRepository {
	return &Pin{
		DB: db,
	}
}

func (p *Pin) CreatePin(pin *models.Pin, boardID int) (*models.Pin, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}

	const query = `
INSERT INTO pins (user_id, title, description, url, image_url, is_private) VALUES (?, ?, ?, ?, ?, ?);
`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(pin.UserID, pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	pinID, err := result.LastInsertId()
	if err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	const query2 = `
INSERT INTO boards_pins (board_id, pin_id) VALUES (?, ?);
`

	stmt, err = tx.Prepare(query2)
	if err != nil {
		return nil, err
	}

	result, err = stmt.Exec(boardID, pinID)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	return pin, nil
}

func (p *Pin) UpdatePin(pin *models.Pin) error {
	return nil
}

func (p *Pin) DeletePin(pinID int) error {
	return nil
}

func (p *Pin) GetPin(pinID int) (*models.Pin, error) {
	const query = `
SELECT
    p.id,
    p.user_id,
    p.title,
    p.description,
    p.url,
    p.image_url,
    p.is_private,
    p.created_at,
    p.updated_at
FROM
    pins p
WHERE
    p.id = ?;
`
	row := p.DB.QueryRow(query, pinID)

	pin := &models.Pin{}
	err := row.Scan(
		&pin.ID,
		&pin.UserID,
		&pin.Title,
		&pin.Description,
		&pin.URL,
		&pin.ImageURL,
		&pin.IsPrivate,
		&pin.CreatedAt,
		&pin.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return pin, nil
}

func (p *Pin) GetPinsByBoardID(boardID int, page int) ([]*models.Pin, error) {
	const query = `
SELECT
    p.id,
    p.user_id,
    p.title,
    p.description,
    p.url,
    p.image_url,
    p.is_private,
    p.created_at,
    p.updated_at
FROM
    pins AS p
    JOIN boards_pins AS bp ON p.id = bp.pin_id
WHERE
	bp.board_id = ?
LIMIT ?
OFFSET ?;
`
	limit := 10
	offset := (page - 1) * limit

	rows, err := p.DB.Query(query, boardID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pins []*models.Pin
	for rows.Next() {
		pin := &models.Pin{}
		err := rows.Scan(
			&pin.ID,
			&pin.UserID,
			&pin.Title,
			&pin.Description,
			&pin.URL,
			&pin.ImageURL,
			&pin.IsPrivate,
			&pin.CreatedAt,
			&pin.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pins, nil
}

func (p *Pin) GetPinsByUserID(userID int) ([]*models.Pin, error) {
	const query = `
SELECT
    p.id,
    p.user_id,
    p.title,
    p.description,
    p.url,
    p.image_url,
    p.is_private,
    p.created_at,
    p.updated_at
FROM
    pins p
WHERE
    p.user_id = ?;
`

	rows, err := p.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pins []*models.Pin
	for rows.Next() {
		pin := &models.Pin{}
		err := rows.Scan(
			&pin.ID,
			&pin.UserID,
			&pin.Title,
			&pin.Description,
			&pin.URL,
			&pin.ImageURL,
			&pin.IsPrivate,
			&pin.CreatedAt,
			&pin.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		pins = append(pins, pin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pins, nil
}
