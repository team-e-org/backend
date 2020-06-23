package helpers

import (
	"errors"
	"testing"
)

const (
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	UNAUTHORIZED          = "UNAUTHORIZED"
	NOT_FOUND             = "NOT_FOUND"
	BAD_REQUEST           = "BAD_REQUEST"
)

func TestInternalServerError(t *testing.T) {
	e := NewInternalServerError(errors.New(INTERNAL_SERVER_ERROR))
	if _, ok := e.(error); !ok {
		t.Fatalf("Error() not implemented")
	}
	if _, ok := e.(AppError); !ok {
		t.Fatalf("AppError() not implemented")
	}
	if e.Error() != INTERNAL_SERVER_ERROR {
		t.Fatalf("Invalid error message")
	}
}

func TestUnauthorized(t *testing.T) {
	e := NewUnauthorized(errors.New(UNAUTHORIZED))
	if _, ok := e.(error); !ok {
		t.Fatalf("Error() not implemented")
	}
	if _, ok := e.(AppError); !ok {
		t.Fatalf("AppError() not implemented")
	}
	if e.Error() != UNAUTHORIZED {
		t.Fatalf("Invalid error message")
	}
}

func TestNotFound(t *testing.T) {
	e := NewNotFound(errors.New(NOT_FOUND))
	if _, ok := e.(error); !ok {
		t.Fatalf("Error() not implemented")
	}
	if _, ok := e.(AppError); !ok {
		t.Fatalf("AppError() not implemented")
	}
	if e.Error() != NOT_FOUND {
		t.Fatalf("Invalid error message")
	}
}

func TestBadRequest(t *testing.T) {
	e := NewBadRequest(errors.New(BAD_REQUEST))
	if _, ok := e.(error); !ok {
		t.Fatalf("Error() not implemented")
	}
	if _, ok := e.(AppError); !ok {
		t.Fatalf("AppError() not implemented")
	}
	if e.Error() != BAD_REQUEST {
		t.Fatalf("Invalid error message")
	}
}
