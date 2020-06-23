package db

import (
	"app/infrastructure"
	"app/mocks"
	"app/repository"
	"database/sql"
)

type S3 repository.FileRepository

// TODO add DataStorageInterface
type DataStorage struct {
	DB         *sql.DB
	Users      repository.UserRepository
	Boards     repository.BoardRepository
	Pins       repository.PinRepository
	Tags       repository.TagRepository
	BoardsPins repository.BoardPinRepository
	AWSS3      repository.FileRepository
}

func NewDataStorage(db *sql.DB, s3 S3) *DataStorage {
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
		AWSS3:      s3,
	}
}

func NewRepositoryMock() *DataStorage {
	users := mocks.NewUserRepository()
	boards := mocks.NewBoardRepository()
	pins := mocks.NewPinRepository()
	tags := mocks.NewTagRepository()
	awsS3 := mocks.NewAWSS3Repository()
	return &DataStorage{
		Users:  users,
		Boards: boards,
		Pins:   pins,
		Tags:   tags,
		AWSS3:  awsS3,
	}
}
