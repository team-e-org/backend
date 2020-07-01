package infrastructure

import (
	"app/logs"
	"app/repository"
	"database/sql"
	"errors"
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

func (bp *BoardPin) DeleteBoardPin(boardID int, pinID int) error {
	const query = `
DELETE FROM boards_pins
WHERE board_id = ? and pin_id = ?;
`

	stmt, err := bp.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(boardID, pinID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	logs.Info("boards pins deleted, boardID: %d, pinID: %d, rowsAffected: %d", boardID, pinID, rowsAffected)

	return nil
}
