package mocks

import "app/repository"

type BoardPinMock struct {
}

func NewBoardPinRepository() repository.BoardPinRepository {
	return &BoardPinMock{}
}

func (b BoardPinMock) CreateBoardPin(boardID int, pinID int) error {
	return nil
}