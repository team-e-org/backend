package mocks

import (
	"app/models/db"
	"app/repository"
)

type BoardMock struct {
	ExpectedBoard *db.Board
}

func NewBoardRepository() repository.BoardRepository {
	return &BoardMock{}
}

func (m *BoardMock) AddBoard(board *db.Board) error {
	m.ExpectedBoard = board
	return nil
}

func (m *BoardMock) GetBoard(boardID int) (*db.Board, error) {
	return m.ExpectedBoard, nil
}
