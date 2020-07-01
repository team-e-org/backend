package mocks

import (
	"app/models"
	"app/repository"
)

type AWSLambdaMock struct {
}

func NewAWSLambda() repository.LambdaRepository {
	return &AWSLambdaMock{}
}

func (l *AWSLambdaMock) AttachTags(pin *models.Pin, tags []string) error {
	return nil
}
