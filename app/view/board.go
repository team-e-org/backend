package view

import "app/models"

type Board struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"isPrivate"`
	IsArchive   bool   `json:"isArchive"`
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
