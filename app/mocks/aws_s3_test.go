package mocks

import "testing"

func TestUploadImage(t *testing.T) {
	s3Mock := NewAWSS3Repository()
	err := s3Mock.UploadImage(nil, "filename", "contentType", -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}
