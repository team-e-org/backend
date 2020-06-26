package usecase

import (
	"app/authz"
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"fmt"
)

func UserBoards(data db.DataStorageInterface, authLayer authz.AuthLayerInterface, userID int, currentUserID int) ([]*models.Board, helpers.AppError) {

	boards, err := data.Boards().GetBoardsByUserID(userID)
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

func UpdateUser(data db.DataStorageInterface, user *models.User, userID int) (*models.User, error) {
	fmt.Printf("user: %v\n", *user)
	fmt.Printf("userID: %d\n", userID)
	u, err := data.Users().GetUser(userID)
	if err != nil {
		logs.Error("An error occurred while getting user data: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	if user.ID != u.ID {
		logs.Error("UserIDs do not match error: %v", err)
		return nil, helpers.NewUnauthorized(err)
	}

	if user.Name != "" {
		u.Name = user.Name
	}
	if user.Email != "" {
		u.Email = user.Email
	}
	if user.HashedPassword != "" {
		u.HashedPassword = user.HashedPassword
	}
	if user.Icon != "" {
		u.Icon = user.Icon
	}

	if err := data.Users().UpdateUser(u); err != nil {
		logs.Error("An error occurred: %v", err)
		return nil, helpers.NewInternalServerError(err)
	}

	return u, nil
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
