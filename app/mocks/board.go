package mocks

import (
	"app/models"
	"app/repository"
)

type BoardMock struct {
	Expectemodelsoard *models.Board
}

func NewBoardRepository() repository.BoardRepository {
	return &BoardMock{}
}

func (m *BoardMock) AddBoard(board *db.Board) error {
	m.ExpectedBoard = board
	return nil
}

func (m *BoardMock) GetBoard(boardID int) (*models.Board, error) {
	return m.Expectemodelsoard, nil
}
