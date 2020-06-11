package repository

import "app/models/db"

type UserRepository interface {
	CreateUser(user *db.User) error
	GetUser(userID int) (*db.User, error)
}

type BoardRepository interface {
	CreateBoard(board *db.Board) error
	GetBoard(boardID int) (*db.Board, error)
}

type PinRepository interface {
	CreatePin(pin *db.Pin) error
	GetPin(pinID int) (*db.Pin, error)
}

type TagRepository interface {
	CreateTag(tag *db.Tag) error
	GetTag(tagID int) (*db.Tag, error)
}
