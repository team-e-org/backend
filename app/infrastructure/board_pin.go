package infrastructure

import (
	"app/logs"
	"app/repository"
	"database/sql"
)

type BoardPin struct {
	DB *sql.DB
}

func NewBoardPinRepository(db *sql.DB) repository.BoardPinRepository {
	return &BoardPin{
		DB: db,
	}
}

func (bp *BoardPin) CreateBoardPin(boardID int, pinID int) error {
	const query = `
INSERT INTO
    boards_pins(board_id, pin_id)
VALUES
    (?, ?)
`

	stmt, err := bp.DB.Prepare(query)
	if err != nil {
		logs.Error("An error occurred: %v", err)
		return err
	}

	_, err = stmt.Exec(boardID, pinID)
	if err != nil {
		logs.Error("An error occurred: %v", err)
		return err
	}

	return nil
}
