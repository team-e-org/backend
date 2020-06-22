package helpers

import (
	"fmt"
)

type AppError interface {
	AppError()
	Error() string
}

type InternalServerError struct {
	err error
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Usecase error: %v\n", e.err)
}

func (e *InternalServerError) AppError() {}

func NewInternalServerError(err error) AppError {
	return &InternalServerError{err: err}
}

type Unauthorized struct {
	err error
}

func (e *Unauthorized) Error() string {
	return fmt.Sprintf("Usecase error: %v\n", e.err)
}

func (e *Unauthorized) AppError() {}

func NewUnauthorized(err error) AppError {
	return &Unauthorized{err: err}
}

type NotFound struct {
	err error
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("Usecase error: %v\n", e.err)
}

func (e *NotFound) AppError() {}

func NewNotFound(err error) AppError {
	return &NotFound{err: err}
}
