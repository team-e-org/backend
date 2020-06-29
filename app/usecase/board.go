package usecase

import (
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"fmt"
)

func CreateBoard(data db.DataStorageInterface, board *models.Board) (*models.Board, helpers.AppError) {

	board, err := data.Boards().CreateBoard(board)
	if err != nil {
		logs.Error("An error occurred while creating board: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	return board, nil
}

func UpdateBoard(data db.DataStorageInterface, board *models.Board) (*models.Board, error) {

	b, err := data.Boards().GetBoard(board.ID)
	if err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, helpers.NewBadRequest(err)
	}

	if b.UserID != board.UserID {
		err := fmt.Errorf("UserIDs do not match error")
		logs.Error("An error occurred: %v", err)
		return nil, helpers.NewUnauthorized(err)
	}

	if err := data.Boards().UpdateBoard(board); err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	return board, nil
}
