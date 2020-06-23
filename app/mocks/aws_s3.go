package mocks

import (
	"app/repository"
	"mime/multipart"
)

type AWSS3Mock struct{}

func NewAWSS3Repository() repository.FileRepository {
	return &AWSS3Mock{}
}

func (m *AWSS3Mock) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (url string, err error) {
	return "https://s3/dummy-s3-url", nil
}
