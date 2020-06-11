package registry

import (
	"app/config"
	"app/db"
	"app/infrastructure"
	"app/mocks"
	"app/repository"
)

type Repository struct {
	Users  repository.UserRepository
	Boards repository.BoardRepository
	Pins   repository.PinRepository
	Tags   repository.TagRepository
}

func NewRepository(c config.DBConfig) (*Repository, error) {
	db, err := db.ConnectToMySql(c.Host, c.Port, c.User, c.Password, c.DBName)
	if err != nil {
		return nil, err
	}
	users := infrastructure.NewUserRepository(db)
	boards := infrastructure.NewBoardRepository(db)
	pins := infrastructure.NewPinRepository(db)
	tags := infrastructure.NewTagRepository(db)
	return &Repository{
		Users:  users,
		Boards: boards,
		Pins:   pins,
		Tags:   tags,
	}, nil
}

func NewRepositoryMock() (*Repository, error) {
	users := mocks.NewUserRepository()
	boards := mocks.NewBoardRepository()
	pins := mocks.NewPinRepository()
	tags := mocks.NewTagRepository()
	return &Repository{
		Users:  users,
		Boards: boards,
		Pins:   pins,
		Tags:   tags,
	}, nil
}
