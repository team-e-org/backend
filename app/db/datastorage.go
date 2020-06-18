package db

import (
	"app/config"
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
	AWSS3      repository.FileRepository
}

func NewDataStorage(db *sql.DB, awsConf *config.AWS) *DataStorage {
	users := infrastructure.NewUserRepository(db)
	boards := infrastructure.NewBoardRepository(db)
	pins := infrastructure.NewPinRepository(db)
	tags := infrastructure.NewTagRepository(db)
	boardsPins := infrastructure.NewBoardPinRepository(db)
	awsS3 := infrastructure.NewAWSS3(awsConf.S3)
	return &DataStorage{
		DB:         db,
		Users:      users,
		Boards:     boards,
		Pins:       pins,
		Tags:       tags,
		BoardsPins: boardsPins,
		AWSS3:      awsS3,
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
