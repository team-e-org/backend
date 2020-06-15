package infrastructure

import (
	"app/models"
	"app/repository"
	"database/sql"
)

type Pin struct {
	DB *sql.DB
}

func NewPinRepository(db *sql.DB) repository.PinRepository {
	return &Pin{
		DB: db,
	}
}

func (u *Pin) CreatePin(pin *models.Pin, boardID int) error {
	return nil
}

func (u *Pin) UpdatePin(pin *models.Pin) error {
	return nil
}

func (u *Pin) DeletePin(pinID int) error {
	return nil
}

func (u *Pin) GetPin(pinID int) (*models.Pin, error) {
	return nil, nil
}

func (u *Pin) GetPinsByBoardID(boardID int) ([]*models.Pin, error) {
	return nil, nil
}

func (u *Pin) GetPinsByUserID(userID int) ([]*models.Pin, error) {
	return nil, nil
}
