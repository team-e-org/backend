package mocks

import (
	"app/repository"
	"errors"
)

type BoardPinMock struct {
	expectBoardID int
	expectPinID   int
}

func NewBoardPinRepository() repository.BoardPinRepository {
	return &BoardPinMock{}
}

func (b *BoardPinMock) CreateBoardPin(boardID int, pinID int) error {
	b.expectBoardID = boardID
	b.expectPinID = pinID

	return nil
}

func (b *BoardPinMock) DeleteBoardPin(boardID int, pinID int) error {
	if !(b.expectBoardID == boardID && b.expectPinID == pinID) {
		return errors.New("boardID and pinID not match")
	}
	return nil
}
