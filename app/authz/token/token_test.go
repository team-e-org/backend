package token

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestNewToken(t *testing.T) {
	id := NewToken()
	_, err := uuid.FromString(id)
	if err != nil {
		t.Fatalf("Invalid uuid")
	}
}
