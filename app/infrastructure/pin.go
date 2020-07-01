package infrastructure

import (
	"app/helpers"
	"app/logs"
	"app/models"
	"app/repository"
	"database/sql"
	"fmt"
)

type Pin struct {
	DB      *sql.DB
	BaseURL string
}

func NewPinRepository(db *sql.DB, baseURL string) repository.PinRepository {
	return &Pin{
		DB:      db,
		BaseURL: baseURL,
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
		return nil, helpers.TryRollback(tx, err)
	}

	result, err := stmt.Exec(pin.UserID, pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	pinID, err := result.LastInsertId()
	if err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	pin.ID = int(pinID)

	const query2 = `
INSERT INTO boards_pins (board_id, pin_id) VALUES (?, ?);
`

	stmt, err = tx.Prepare(query2)
	if err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	result, err = stmt.Exec(boardID, pinID)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, helpers.TryRollback(tx, err)
	}

	logs.Info("New pin created, id: %v", int(pinID))

	pin.ImageURL = fmt.Sprintf("%s/%s", p.BaseURL, pin.ImageURL)

	return pin, nil
}

func (p *Pin) UpdatePin(pin *models.Pin) error {
	const query = `
UPDATE pins SET title = ?, description = ?, url = ?, image_url = ?, is_private = ? WHERE id = ?;
`

	stmt, err := p.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(pin.Title, pin.Description, pin.URL, pin.ImageURL, pin.IsPrivate, pin.ID)
	if err := helpers.CheckDBExecError(result, err); err != nil {
		return err
	}

	logs.Info("Pin updated, id: %v", pin.ID)

	pin.ImageURL = fmt.Sprintf("%s/%s", p.BaseURL, pin.ImageURL)

	return nil
}

func (p *Pin) DeletePin(pinID int) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}

	const query = `
DELETE FROM pins WHERE id = ?;
`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return helpers.TryRollback(tx, err)
	}

	result, err := stmt.Exec(pinID)
	if err := helpers.CheckDBExecError(result, err); err != nil {
		return helpers.TryRollback(tx, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return helpers.TryRollback(tx, err)
	}

	const query2 = `
DELETE FROM boards_pins WHERE pin_id = ?;
`

	stmt, err = tx.Prepare(query2)
	if err != nil {
		return helpers.TryRollback(tx, err)
	}

	result, err = stmt.Exec(pinID)
	if err := helpers.CheckDBExecError(result, err); err != nil {
		return helpers.TryRollback(tx, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return helpers.TryRollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return helpers.TryRollback(tx, err)
	}

	logs.Info("Pin deleted, id: %v", pinID)

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

	stmt, err := p.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(pinID)

	pin := &models.Pin{}
	err = row.Scan(
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

	logs.Info("Pin got, id: %v", pinID)

	pin.ImageURL = fmt.Sprintf("%s/%s", p.BaseURL, pin.ImageURL)

	return pin, nil
}

func (p *Pin) GetPins(page int) ([]*models.Pin, error) {
	query := `
SELECT
    p.id,
    p.user_id,
    p.title,
    p.description,
    p.url,
    p.image_url,
    p.created_at,
    p.updated_at
FROM
    pins AS p
WHERE is_private = 0
ORDER BY created_at DESC
LIMIT ? OFFSET ?;
`
	limit := 10
	offset := (page - 1) * limit

	stmt, err := p.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit, offset)
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
			&pin.CreatedAt,
			&pin.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		pin.ImageURL = fmt.Sprintf("%s/%s", p.BaseURL, pin.ImageURL)
		pins = append(pins, pin)
	}

	logs.Info("Pins got, page: %v", page)

	return pins, nil
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
LIMIT ? OFFSET ?;
`
	limit := 10
	offset := (page - 1) * limit

	stmt, err := p.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(boardID, limit, offset)
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
		pin.ImageURL = fmt.Sprintf("%s/%s", p.S3URL, pin.ImageURL)
		pins = append(pins, pin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	logs.Info("Pins got, boardID: %v, page: %v", boardID, page)

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

	stmt, err := p.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
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
		pin.ImageURL = fmt.Sprintf("%s/%s", p.S3URL, pin.ImageURL)
		pins = append(pins, pin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	logs.Info("Pins got, userID: %v", userID)

	return pins, nil
}
