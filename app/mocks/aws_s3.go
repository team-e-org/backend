package mocks

import (
	"app/repository"
	"mime/multipart"
)

type AWSS3Mock struct {
	ExpectedURL string
	URL         string
}

func NewAWSS3Repository() repository.FileRepository {
	return &AWSS3Mock{ExpectedURL: "dummy-s3-url", URL: "https://s3"}
}

func (m *AWSS3Mock) UploadImage(file multipart.File, fileHeader *multipart.FileHeader, userID int) (url string, err error) {
	return m.ExpectedURL, nil
}

func (m *AWSS3Mock) GetURL() string {
	return m.URL
}
