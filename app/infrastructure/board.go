package infrastructure

import (
	"app/helpers"
	"app/models"
	"app/repository"
	"database/sql"
)

type Board struct {
	DB *sql.DB
}

func NewBoardRepository(db *sql.DB) repository.BoardRepository {
	return &Board{
		DB: db,
	}
}

func (b *Board) CreateBoard(board *models.Board) (*models.Board, error) {
	const query = `
INSERT INTO boards (user_id, name, description, is_private) VALUES (?, ?, ?, ?)
`

	stmt, err := b.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(
		board.UserID,
		board.Name,
		board.Description,
		board.IsPrivate)
	err = helpers.CheckDBExecError(result, err)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	board.ID = int(id)

	return board, nil
}

func (b *Board) UpdateBoard(board *models.Board) error {
	const query = `
UPDATE boards SET name = ?, description = ?, is_private = ?;
`

	stmt, err := b.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(board.Name, board.Description, board.IsPrivate)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return err
	}

	return nil
}

func (b *Board) DeleteBoard(boardID int) error {
	const query = `
DELETE FROM boards WHERE id = ?;
`

	stmt, err := b.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(boardID)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return err
	}

	return nil
}

func (b *Board) GetBoard(boardID int) (*models.Board, error) {
	const query = `
SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at FROM boards b WHERE b.id = ?;
`

	stmt, err := b.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(boardID)

	board := &models.Board{}
	err = row.Scan(
		&board.ID,
		&board.UserID,
		&board.Name,
		&board.Description,
		&board.IsPrivate,
		&board.IsArchive,
		&board.CreatedAt,
		&board.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return board, nil
}

func (b *Board) GetBoardsByUserID(userID int) ([]*models.Board, error) {
	const query = `
SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at FROM boards b WHERE b.user_id = ?;
`

	stmt, err := b.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boards := make([]*models.Board, 0)
	for rows.Next() {
		board := &models.Board{}
		err := rows.Scan(
			&board.ID,
			&board.UserID,
			&board.Name,
			&board.Description,
			&board.IsPrivate,
			&board.IsArchive,
			&board.CreatedAt,
			&board.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}
