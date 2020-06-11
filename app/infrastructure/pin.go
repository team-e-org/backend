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

func (u *Pin) CreatePin(pin *models.Pin) error {
	return nil
}

func (u *Pin) GetPin(pinID int) (*models.Pin, error) {
	return nil, nil
}
