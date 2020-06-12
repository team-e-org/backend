package view

import "app/models"

type Board struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
	IsArchive   bool   `json:"is_archive"`
}

func NewBoard(board *models.Board) *Board {
	b := &Board{
		ID:          board.ID,
		UserID:      board.UserID,
		Name:        board.Name,
		Description: board.Description,
		IsPrivate:   board.IsPrivate,
		IsArchive:   board.IsArchive,
	}

	return b
}
