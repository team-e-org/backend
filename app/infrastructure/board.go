package infrastructure

import (
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

func (u *Board) CreateBoard(board *models.Board) error {
	return nil
}

func (u *Board) UpdateBoard(board *models.Board) error {
	return nil
}

func (u *Board) DeleteBoard(boardID int) error {
	return nil
}

func (b *Board) GetBoard(boardID int) (*models.Board, error) {
	const query = `
SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at
FROM boards b
WHERE b.id = ?;
`
	row := b.DB.QueryRow(query, boardID)

	board := &models.Board{}
	err := row.Scan(
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
SELECT b.id, b.user_id, b.name, b.description, b.is_private, b.is_archive, b.created_at, b.updated_at
FROM boards b
WHERE b.user_id = ?;
`

	rows, err := b.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boards []*models.Board
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
