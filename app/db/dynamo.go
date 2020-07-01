package db

import (
	"app/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func ConnectToDynamo(config config.Dynamo) *dynamo.DB {
	return dynamo.New(session.New(), &aws.Config{Region: aws.String(config.Region)})
}
