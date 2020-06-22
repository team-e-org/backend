package mocks

import "mime/multipart"

type AWSS3Mock struct {
	ExpectedURL string
}

func (m *AWSS3Mock) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (url string, err error) {
	return m.ExpectedURL, nil
}
