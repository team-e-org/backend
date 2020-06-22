package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCheckDBExecError(t *testing.T) {
	err := CheckDBExecError(sqlmock.NewResult(0, 1), errors.New("An error"))
	if err == nil {
		t.Fatalf("An error should occur")
	}

	err = CheckDBExecError(sqlmock.NewResult(0, 2), nil)
	if err == nil {
		t.Fatalf("An error should occur")
	}

	err = CheckDBExecError(sqlmock.NewResult(0, 1), nil)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("An error occurred")
	}
	if password == hashedPassword {
		t.Fatalf("The password is not hashed")
	}
}
