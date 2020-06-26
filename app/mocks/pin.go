package mocks

import (
	"app/models"
	"app/repository"
	"errors"
)

type PinMock struct {
	ExpectedPins   []*models.Pin
	BoardPinMapper map[int][]int // map[boardID][]pinID
}

func NewPinRepository() repository.PinRepository {
	pins := make([]*models.Pin, 0)
	mapper := make(map[int][]int)
	return &PinMock{
		ExpectedPins:   pins,
		BoardPinMapper: mapper,
	}
}

func (m *PinMock) CreatePin(pin *models.Pin, boardID int) (*models.Pin, error) {
	m.ExpectedPins = append(m.ExpectedPins, pin)
	m.BoardPinMapper[boardID] = append(m.BoardPinMapper[boardID], pin.ID)
	return pin, nil
}

func (m *PinMock) UpdatePin(pin *models.Pin) error {
	for i, p := range m.ExpectedPins {
		if p.ID == pin.ID {
			m.ExpectedPins[i] = pin
			return nil
		}
	}
	return noPinError()
}

func (m *PinMock) DeletePin(pinID int) error {
	for i, p := range m.ExpectedPins {
		if p.ID == pinID {
			m.ExpectedPins = append(m.ExpectedPins[:i], m.ExpectedPins[i+1:]...)
			return nil
		}
	}
	return noPinError()
}

func (m *PinMock) GetPin(pinID int) (*models.Pin, error) {
	for _, p := range m.ExpectedPins {
		if p.ID == pinID {
			return p, nil
		}
	}
	return nil, noPinError()
}

func (m *PinMock) GetPinsByBoardID(boardID int, page int) ([]*models.Pin, error) {
	if _, ok := m.BoardPinMapper[boardID]; !ok {
		return nil, noBoardError()
	}
	pins := make([]*models.Pin, 0, len(m.BoardPinMapper))
	for _, id := range m.BoardPinMapper[boardID] {
		for _, p := range m.ExpectedPins {
			if p.ID == id {
				pins = append(pins, p)
			}
		}
	}
	return pins, nil
}

func (m *PinMock) GetPinsByUserID(userID int) ([]*models.Pin, error) {
	pins := make([]*models.Pin, 0, len(m.BoardPinMapper))
	for _, p := range m.ExpectedPins {
		if *p.UserID == userID {
			pins = append(pins, p)
		}
	}
	return pins, nil
}

func (m *PinMock) GetPins(page int) ([]*models.Pin, error) {
	return nil, nil
}

func noPinError() error {
	return errors.New("The pin does not exist")
}
