package repository

import "app/models/db"

type UserRepository interface {
	AddUser(user *db.User) error
	GetUser(userID int) (*db.User, error)
}

type BoardRepository interface {
	AddBoard(board *db.Board) error
	GetBoard(boardID int) (*db.Board, error)
}

type PinRepository interface {
	AddPin(pin *db.Pin) error
	GetPin(pinID int) (*db.Pin, error)
}

type TagRepository interface {
	AddTag(tag *db.Tag) error
	GetTag(tagID int) (*db.Tag, error)
}
