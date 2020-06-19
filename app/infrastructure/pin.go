package infrastructure

import (
	"app/helpers"
	"app/logs"
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
	const query = `
INSERT INTO pins (
	user_id,
    title,
    description,
    url,
    is_private,
    image_url,
    created_at,
    updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`
	stmt, err := p.DB.Prepare(query)
	if err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, err
	}

	result, err := stmt.Exec(
		pin.UserID,
		pin.Title,
		pin.Description,
		pin.URL,
		pin.IsPrivate,
		pin.ImageURL,
		pin.CreatedAt,
		pin.UpdatedAt)
	err = helpers.CheckDBExecError(result, err)
	if err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, err
	}
	pin.ID = int(id)

	return pin, nil
}

func (u *Pin) UpdatePin(pin *models.Pin) error {
	return nil
}

func (u *Pin) DeletePin(pinID int) error {
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
