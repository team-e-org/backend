package helpers

import (
	"database/sql"
	"errors"
)

func CheckDBExecError(result sql.Result, err error) error {
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("Affected more than 1 rows")
	}

	return nil
}
