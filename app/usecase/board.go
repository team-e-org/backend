package usecase

import (
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"app/view"
)

func CreateBoard(data db.DataStorageInterface, requestBoard *view.Board, userID int) (*models.Board, helpers.AppError) {

	board := &models.Board{
		UserID:      userID,
		Name:        requestBoard.Name,
		Description: &requestBoard.Description,
		IsPrivate:   requestBoard.IsPrivate,
	}
	board, err := data.Boards().CreateBoard(board)
	if err != nil {
		logs.Error("An error occurred while creating board: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	return board, nil
}

func UpdateBoard(data db.DataStorageInterface, board *models.Board) (*models.Board, error) {

	if err := data.Boards().UpdateBoard(board); err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	return board, nil
}
