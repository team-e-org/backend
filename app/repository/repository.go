package repository

import "app/models"

type UserRepository interface {
	CreateUser(name string, email string, icon string, password string) (*models.User, error)
	GetUser(userID int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type BoardRepository interface {
	CreateBoard(board *models.Board) error
	GetBoard(boardID int) (*models.Board, error)
}

type PinRepository interface {
	CreatePin(pin *models.Pin) error
	GetPin(pinID int) (*models.Pin, error)
}

type TagRepository interface {
	CreateTag(tag *models.Tag) error
	GetTag(tagID int) (*models.Tag, error)
}
