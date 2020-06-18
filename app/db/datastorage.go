package db

import (
	"app/infrastructure"
	"app/mocks"
	"app/repository"
	"database/sql"
)

type DataStorage struct {
	DB         *sql.DB
	Users      repository.UserRepository
	Boards     repository.BoardRepository
	Pins       repository.PinRepository
	Tags       repository.TagRepository
	BoardsPins repository.BoardPinRepository
}

func NewDataStorage(db *sql.DB) *DataStorage {
	users := infrastructure.NewUserRepository(db)
	boards := infrastructure.NewBoardRepository(db)
	pins := infrastructure.NewPinRepository(db)
	tags := infrastructure.NewTagRepository(db)
	boardsPins := infrastructure.NewBoardPinRepository(db)
	return &DataStorage{
		DB:         db,
		Users:      users,
		Boards:     boards,
		Pins:       pins,
		Tags:       tags,
		BoardsPins: boardsPins,
	}
}

func NewRepositoryMock() *DataStorage {
	users := mocks.NewUserRepository()
	boards := mocks.NewBoardRepository()
	pins := mocks.NewPinRepository()
	tags := mocks.NewTagRepository()
	return &DataStorage{
		Users:  users,
		Boards: boards,
		Pins:   pins,
		Tags:   tags,
	}
}
