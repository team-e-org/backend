package usecase

import (
	"app/authz"
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
)

func UserBoards(data *db.DataStorage, authLayer authz.AuthLayerInterface, userID int, currentUserID int) ([]*models.Board, helpers.AppError) {

	boards, err := data.Boards.GetBoardsByUserID(userID)
	if err != nil {
		logs.Error("An error occurred while getting user's boards: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	boards = removePrivateBoards(boards, currentUserID)

	if len(boards) == 0 {
		logs.Error("Board not found for userID: %d", userID)
		return nil, helpers.NewNotFound(err)
	}

	return boards, nil
}

func removePrivateBoards(boards []*models.Board, userID int) []*models.Board {
	res := make([]*models.Board, 0)
	for _, b := range boards {
		if b.IsPrivate && b.UserID != userID {
			continue
		}
		res = append(res, b)
	}

	return res
}
