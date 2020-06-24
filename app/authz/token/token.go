package token

import uuid "github.com/satori/go.uuid"

func NewToken() string {
	id := uuid.NewV4()
	return id.String()
}
