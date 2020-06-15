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

func (u *Board) GetBoard(boardID int) (*models.Board, error) {
	return nil, nil
}
