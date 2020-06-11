package infrastructure

import (
	"app/models/db"
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

func (u *Board) CreateBoard(board *db.Board) error {
	return nil
}

func (u *Board) GetBoard(boardID int) (*db.Board, error) {
	return nil, nil
}
