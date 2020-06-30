package repository

import (
	"app/models"
	"mime/multipart"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID int) error
	GetUser(userID int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type BoardRepository interface {
	CreateBoard(board *models.Board) (*models.Board, error)
	UpdateBoard(board *models.Board) error
	DeleteBoard(boardID int) error
	GetBoard(boardID int) (*models.Board, error)
	GetBoardsByUserID(userID int) ([]*models.Board, error)
}

type PinRepository interface {
	CreatePin(pin *models.Pin, boardID int) (*models.Pin, error)
	UpdatePin(pin *models.Pin) error
	DeletePin(pinID int) error
	GetPin(pinID int) (*models.Pin, error)
	GetPinsByBoardID(boardID int, page int) ([]*models.Pin, error)
	GetPinsByUserID(userID int) ([]*models.Pin, error)
	GetPins(page int) ([]*models.Pin, error)
}

type TagRepository interface {
	CreateTag(tag *models.Tag) (*models.Tag, error)
	GetTag(tagID int) (*models.Tag, error)
	AttachTagToPin(tagID int, pinID int) error
	GetTagsByPinID(pinID int) ([]*models.Tag, error)
}

type BoardPinRepository interface {
	CreateBoardPin(boardID int, pinID int) error
}

type FileRepository interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader, userID int) (url string, err error)
	GetURL() string
}

type LambdaRepository interface {
	AttachTags(pin *models.Pin, tags []string) error
}
