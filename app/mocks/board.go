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

func (m *BoardMock) CreateBoard(board *models.Board) error {
	m.Expectemodelsoard = board
	return nil
}

func (m *BoardMock) GetBoard(boardID int) (*models.Board, error) {
	return m.Expectemodelsoard, nil
}
