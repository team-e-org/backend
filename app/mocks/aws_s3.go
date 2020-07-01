package mocks

import (
	"app/repository"
	"mime/multipart"
)

type AWSS3Mock struct {
	FileName string
	URL      string
}

func NewAWSS3Repository() repository.FileRepository {
	return &AWSS3Mock{FileName: "dummy-s3-url", URL: "https://s3"}
}

func (m *AWSS3Mock) UploadImage(file multipart.File, fileName string, contentType string, userID int) error {
	return nil
}

func (m *AWSS3Mock) GetBaseURL() string {
	return m.BaseURL
}

func (m *AWSS3Mock) GetPinFolder() string {
	return "pins"
}

func (m *AWSS3Mock) CreateFileName(userID int, fileExt string) string {
	return m.FileName
}
