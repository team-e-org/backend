package mocks

import "testing"

func TestUploadImage(t *testing.T) {
	s3Mock := &AWSS3Mock{ExpectedURL: "https://s3/tekitouna-gazou.jpg"}
	got, err := s3Mock.UploadImage(nil, nil)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
	if s3Mock.ExpectedURL != got {
		t.Fatalf("want: %v, got %v", s3Mock.ExpectedURL, got)
	}
}
