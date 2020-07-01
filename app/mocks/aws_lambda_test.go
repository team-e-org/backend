package mocks

import (
	"context"
	"testing"
)

func TestAWSLambdaMock_AttachTags(t *testing.T) {
	mock := NewAWSLambda()
	ctx := context.TODO()
	err := mock.AttachTagsWithContext(ctx, nil, nil)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}
