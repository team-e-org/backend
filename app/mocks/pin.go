package mocks

import (
	"app/models"
	"app/repository"
)

type PinMock struct {
	ExpectedPin *models.Pin
}

func NewPinRepository() repository.PinRepository {
	return &PinMock{}
}

func (m *PinMock) CreatePin(pin *models.Pin) error {
	m.ExpectedPin = pin
	return nil
}

func (m *PinMock) GetPin(pinID int) (*models.Pin, error) {
	return m.ExpectedPin, nil
}
