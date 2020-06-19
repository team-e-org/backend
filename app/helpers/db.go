package helpers

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
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

func TryRollback(tx *sql.Tx, err error) error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	return err
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
