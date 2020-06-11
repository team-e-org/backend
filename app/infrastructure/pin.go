package infrastructure

import (
	"app/models/db"
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

func (u *Pin) CreatePin(pin *db.Pin) error {
	return nil
}

func (u *Pin) GetPin(pinID int) (*db.Pin, error) {
	return nil, nil
}
