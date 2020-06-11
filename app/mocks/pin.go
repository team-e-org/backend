package mocks

import (
	"app/models/db"
	"app/repository"
)

type PinMock struct {
	ExpectedPin *db.Pin
}

func NewPinRepository() repository.PinRepository {
	return &PinMock{}
}

func (m *PinMock) CreatePin(pin *db.Pin) error {
	m.ExpectedPin = pin
	return nil
}

func (m *PinMock) GetPin(pinID int) (*db.Pin, error) {
	return m.ExpectedPin, nil
}
