package helpers

type AppError interface {
	AppError()
	Error() string
}

type InternalServerError struct {
	err error
}

func (e *InternalServerError) Error() string {
	return e.err.Error()
}

func (e *InternalServerError) AppError() {}

func NewInternalServerError(err error) AppError {
	return &InternalServerError{err: err}
}

type Unauthorized struct {
	err error
}

func (e *Unauthorized) Error() string {
	return e.err.Error()
}

func (e *Unauthorized) AppError() {}

func NewUnauthorized(err error) AppError {
	return &Unauthorized{err: err}
}

type NotFound struct {
	err error
}

func (e *NotFound) Error() string {
	return e.err.Error()
}

func (e *NotFound) AppError() {}

func NewNotFound(err error) AppError {
	return &NotFound{err: err}
}

type BadRequest struct {
	err error
}

func (e *BadRequest) Error() string {
	return e.err.Error()
}

func (e *BadRequest) AppError() {}

func NewBadRequest(err error) AppError {
	return &BadRequest{err: err}
}
