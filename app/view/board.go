package view

import "app/models"

type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"isPrivate"`
	IsArchive   bool   `json:"isArchive"`
}

func NewBoard(board *models.Board) *Board {
	b := &Board{
		ID:          board.ID,
		Name:        board.Name,
		Description: *board.Description,
		IsPrivate:   board.IsPrivate,
		IsArchive:   board.IsArchive,
	}

	return b
}

func NewBoards(boards []*models.Board) []*Board {
	bs := make([]*Board, 0, len(boards))

	for _, board := range boards {
		bs = append(bs, NewBoard(board))
	}

	return bs
}

func NewBoardModel(board *Board, userID int) *models.Board {
	b := &models.Board{
		ID:          board.ID,
		UserID:      userID,
		Name:        board.Name,
		Description: &board.Description,
		IsPrivate:   board.IsPrivate,
		IsArchive:   board.IsArchive,
	}

	return b
}
