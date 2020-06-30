package infrastructure

import (
	"app/config"
	"app/logs"
	"app/repository"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/satori/go.uuid"
)

type AWSS3 struct {
	Config   config.S3
	Uploader *s3manager.Uploader
}

func NewAWSS3(c config.S3) repository.FileRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(c.Region),
		},
	}))

	logs.Info("Connected to AWSS3 region: %s, bucket: %s, pin-folder: %s", c.Region, c.Bucket, c.PinFolder)
	return &AWSS3{
		Config:   c,
		Uploader: s3manager.NewUploader(sess),
	}
}

type AWSS3Mock struct{}

func (a *AWSS3Mock) UploadImage(file multipart.File, fileHeader *multipart.FileHeader, userID int) (string, error) {
	return "", nil
}

func (a *AWSS3Mock) GetBaseURL() string {
	return ""
}

func NewAWSS3Mock() repository.FileRepository {
	return &AWSS3Mock{}
}

func (a *AWSS3) UploadImage(file multipart.File, fileHeader *multipart.FileHeader, userID int) (url string, err error) {
	var contentType string
	fileExt := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%s/%d/%s%s", a.Config.PinFolder, userID, uuid.NewV4().String(), fileExt)

	logs.Info("File uploaded to %s", fileName)

	switch fileExt {
	case ".jpg":
		contentType = "image/jpeg"
	case ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	default:
		return "", fmt.Errorf("this extension is invalid, %v", fileExt)
	}

	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        file,
		Bucket:      aws.String(a.Config.Bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(fileName),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	logs.Info("File location is %s", result.Location)

	return fileName, nil
}

func (a *AWSS3) GetBaseURL() string {
	return a.Config.BaseURL
}
