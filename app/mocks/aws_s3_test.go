package mocks

import "testing"

func TestUploadImage(t *testing.T) {
	s3Mock := NewAWSS3Repository()
	err := s3Mock.UploadImage(nil, "filename", "contentType", -1)
	if err != nil {
		t.Fatalf("An error occurred: %v", err)
	}
}

func TestGetBaseURL(t *testing.T) {
	s3Mock := NewAWSS3Repository()
	want := "https://s3"
	got := s3Mock.GetBaseURL()

	if want != got {
		t.Fatalf("want: %v, got %v", want, got)
	}
}

func TestGetPinFolder(t *testing.T) {
	s3Mock := NewAWSS3Repository()
	want := "pins"
	got := s3Mock.GetPinFolder()

	if want != got {
		t.Fatalf("want: %v, got %v", want, got)
	}
}

func TestCreateFileName(t *testing.T) {
	s3Mock := NewAWSS3Repository()
	want := "dummy-s3-url"
	got := s3Mock.CreateFileName(0, "")

	if want != got {
		t.Fatalf("want: %v, got %v", want, got)
	}
}
