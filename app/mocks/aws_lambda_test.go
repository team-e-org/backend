package mocks

import "testing"

func TestAWSLambdaMock_AttachTags(t *testing.T) {
	mock := NewAWSLambda()
	err := mock.AttachTags(nil, nil)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}
