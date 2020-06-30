package infrastructure

import (
	"app/config"
	"app/logs"
	"app/repository"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func (a *AWSS3) UploadImage(file multipart.File, fileName string, contentType string, userID int) error {
	logs.Info("File uploaded to %s", fileName)

	result, err := a.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        file,
		Bucket:      aws.String(a.Config.Bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(fileName),
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	logs.Info("File location is %s", result.Location)

	return nil
}

func (a *AWSS3) GetBaseURL() string {
	return a.Config.BaseURL
}

func (a *AWSS3) GetPinFolder() string {
	return a.Config.PinFolder
}
