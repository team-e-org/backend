package mocks

import (
	"app/repository"
	"mime/multipart"
)

type AWSS3Mock struct {
	ExpectedURL string
}

func NewAWSS3Repository() repository.FileRepository {
	return &AWSS3Mock{ExpectedURL: "https://s3/dummy-s3-url"}
}

func (m *AWSS3Mock) UploadImage(file multipart.File, fileHeader *multipart.FileHeader, userID int) (url string, err error) {
	return m.ExpectedURL, nil
}
