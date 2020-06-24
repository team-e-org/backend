package db

import (
	"app/infrastructure"
	"app/mocks"
	"app/repository"
	"database/sql"
)

type S3 repository.FileRepository

type DataStorageInterface interface {
	DB() *sql.DB
	Users() repository.UserRepository
	Boards() repository.BoardRepository
	Pins() repository.PinRepository
	Tags() repository.TagRepository
	BoardsPins() repository.BoardPinRepository
	AWSS3() repository.FileRepository
}

type DataStorage struct {
	db         *sql.DB
	users      repository.UserRepository
	boards     repository.BoardRepository
	pins       repository.PinRepository
	tags       repository.TagRepository
	boardsPins repository.BoardPinRepository
	awss3      repository.FileRepository
}

func (d *DataStorage) DB() *sql.DB                               { return d.db }
func (d *DataStorage) Users() repository.UserRepository          { return d.users }
func (d *DataStorage) Boards() repository.BoardRepository        { return d.boards }
func (d *DataStorage) Pins() repository.PinRepository            { return d.pins }
func (d *DataStorage) Tags() repository.TagRepository            { return d.tags }
func (d *DataStorage) BoardsPins() repository.BoardPinRepository { return d.boardsPins }
func (d *DataStorage) AWSS3() repository.FileRepository          { return d.awss3 }

func NewDataStorage(db *sql.DB, s3 S3) DataStorageInterface {
	users := infrastructure.NewUserRepository(db)
	boards := infrastructure.NewBoardRepository(db)
	pins := infrastructure.NewPinRepository(db)
	tags := infrastructure.NewTagRepository(db)
	boardsPins := infrastructure.NewBoardPinRepository(db)
	return &DataStorage{
		db:         db,
		users:      users,
		boards:     boards,
		pins:       pins,
		tags:       tags,
		boardsPins: boardsPins,
		awss3:      s3,
	}
}

func NewRepositoryMock() DataStorageInterface {
	users := mocks.NewUserRepository()
	boards := mocks.NewBoardRepository()
	pins := mocks.NewPinRepository()
	tags := mocks.NewTagRepository()
	awsS3 := mocks.NewAWSS3Repository()
	return &DataStorage{
		users:  users,
		boards: boards,
		pins:   pins,
		tags:   tags,
		awss3:  awsS3,
	}
}
