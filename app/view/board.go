package view

import "app/models"

type Board struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsPrivate   bool   `json:"isPrivate"`
	IsArchive   bool   `json:"isArchive"`
}

func NewBoard(board *models.Board) *Board {
	b := &Board{
		ID:          board.ID,
		UserID:      board.UserID,
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
