package mocks

import (
	"app/models"
	"app/repository"
	"errors"
)

type BoardMock struct {
	Boards []*models.Board
}

func NewBoardRepository() repository.BoardRepository {
	boards := make([]*models.Board, 0)
	return &BoardMock{
		Boards: boards,
	}
}

func (m *BoardMock) CreateBoard(board *models.Board) (*models.Board, error) {
	m.Boards = append(m.Boards, board)
	return m.Boards[len(m.Boards)-1], nil
}

func (m *BoardMock) UpdateBoard(board *models.Board) error {
	for i, b := range m.Boards {
		if b.ID == board.ID {
			m.Boards[i] = board
			return nil
		}
	}
	return noBoardError()
}

func (m *BoardMock) DeleteBoard(boardID int) error {
	for i, b := range m.Boards {
		if b.ID == boardID {
			m.Boards = append(m.Boards[:i], m.Boards[i+1:]...)
			return nil
		}
	}
	return noBoardError()
}

func (m *BoardMock) GetBoard(boardID int) (*models.Board, error) {
	for _, b := range m.Boards {
		if b.ID == boardID {
			return b, nil
		}
	}
	return nil, noBoardError()
}

func (m *BoardMock) GetBoardsByUserID(userID int) ([]*models.Board, error) {
	boards := make([]*models.Board, 0)
	for _, b := range m.Boards {
		if b.UserID == userID {
			boards = append(boards, b)
		}
	}
	return boards, nil
}

func noBoardError() error {
	return errors.New("An error occurred, the board does not exist")
}
