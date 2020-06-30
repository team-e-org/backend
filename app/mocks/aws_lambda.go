package mocks

import (
	"app/models"
	"app/repository"
	"context"
)

type AWSLambdaMock struct {
}

func NewAWSLambda() repository.LambdaRepository {
	return &AWSLambdaMock{}
}

func (l *AWSLambdaMock) AttachTagsWithContext(ctx context.Context, pin *models.Pin, tags []string) error {
	return nil
}
