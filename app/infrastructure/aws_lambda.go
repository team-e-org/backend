package infrastructure

import (
	"app/config"
	"app/logs"
	"app/models"
	"app/repository"
	"app/view"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type AWSLambda struct {
	Config config.Lambda
	svc    *lambda.Lambda
}

func NewAWSLambda(c config.Lambda) repository.LambdaRepository {
	return &AWSLambda{
		Config: c,
		svc:    lambda.New(session.New(), aws.NewConfig().WithRegion(c.Region)),
	}
}

func (l *AWSLambda) AttachTags(pin *models.Pin, tags []string) error {
	lambdaPayload := view.AttachtagsLambdaPayload{
		Pin:  view.NewLambdaPin(pin),
		Tags: tags,
	}

	lambdaPayloadBytes, err := json.Marshal(lambdaPayload)
	if err != nil {
		return err
	}

	logs.Info(string(lambdaPayloadBytes))

	input := &lambda.InvokeInput{
		FunctionName:   aws.String(l.Config.FunctionARN),
		Payload:        lambdaPayloadBytes,
		InvocationType: aws.String(l.Config.InvocationType),
	}

	_, err = l.svc.Invoke(input)
	return err
}
