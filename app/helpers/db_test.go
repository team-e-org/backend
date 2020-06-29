package helpers

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCheckDBExecError(t *testing.T) {
	err := CheckDBExecError(sqlmock.NewResult(0, 1), nil)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	err = CheckDBExecError(sqlmock.NewResult(0, 2), nil)
	if err == nil {
		t.Fatalf("An error should occur")
	}

	err = CheckDBExecError(sqlmock.NewResult(0, 1), errors.New("some error"))
	if err == nil {
		t.Fatalf("An error should occur")
	}

	res := sqlmock.NewErrorResult(errors.New("result error"))
	err = CheckDBExecError(res, nil)
	if err == nil {
		t.Fatalf("An error should occur")
	}
}

func TestHashPassword(t *testing.T) {
	password := "password"
	s, err := HashPassword(password)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}

	if s == password {
		t.Fatalf("String is not hashed")
	}

	if s2, err := HashPassword(password); err != nil && s == s2 {
		t.Fatalf("An error occurred or invalid hashed password")
	}

}
