package db

import (
	"app/config"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/satori/go.uuid"
)

type AwsS3 struct {
	Config   config.S3
	Uploader *s3manager.Uploader
}

func NewAwsS3(c config.S3) *AwsS3 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     c.AccessKeyID,
				SecretAccessKey: c.SecretAccessKey,
			}),
			Region: aws.String(c.Region),
		},
	}))

	return &AwsS3{
		Config:   c,
		Uploader: s3manager.NewUploader(sess),
	}
}

func (a *AwsS3) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (url string, err error) {
	var contentType string
	fileExt := filepath.Ext(fileHeader.Filename)
	fileName := a.Config.PinFolder + uuid.NewV4().String() + fileExt

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

	return result.Location, nil
}
