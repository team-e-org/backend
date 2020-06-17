package mocks

import (
	"app/models"
	"app/repository"
	"errors"
)

type BoardMock struct {
	ExpectedBoard *models.Board
}

func NewBoardRepository() repository.BoardRepository {
	return &BoardMock{}
}

func (m *BoardMock) CreateBoard(board *models.Board) (*models.Board, error) {
	m.ExpectedBoard = board
	return m.ExpectedBoard, nil
}

func (m *BoardMock) UpdateBoard(board *models.Board) error {
	if m.ExpectedBoard == nil {
		return noBoardError()
	}
	m.ExpectedBoard = board
	return nil
}

func (m *BoardMock) DeleteBoard(boardID int) error {
	if m.ExpectedBoard.ID != boardID {
		return noBoardError()
	}
	return nil
}

func (m *BoardMock) GetBoard(boardID int) (*models.Board, error) {
	if m.ExpectedBoard.ID != boardID {
		return nil, noBoardError()
	}
	return m.ExpectedBoard, nil
}

func (m *BoardMock) GetBoardsByUserID(userID int) ([]*models.Board, error) {
	return []*models.Board{m.ExpectedBoard}, nil
}

func noBoardError() error {
	return errors.New("An error occurred, the board does not exist")
}
